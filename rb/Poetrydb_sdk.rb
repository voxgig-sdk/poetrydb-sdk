# Poetrydb SDK

require_relative 'utility/struct/voxgig_struct'
require_relative 'core/utility_type'
require_relative 'core/spec'
require_relative 'core/helpers'

# Load utility registration
require_relative 'utility/register'

# Load config and features
require_relative 'config'
require_relative 'feature/base_feature'
require_relative 'features'

# Load typed models (Struct value objects).
require_relative 'Poetrydb_types'


class PoetrydbSDK
  attr_accessor :mode, :features, :options

  def initialize(options = {})
    @mode = "live"
    @features = []
    @options = nil

    utility = PoetrydbUtility.new
    @_utility = utility

    config = PoetrydbConfig.make_config

    @_rootctx = utility.make_context.call({
      "client" => self,
      "utility" => utility,
      "config" => config,
      "options" => options || {},
      "shared" => {},
    }, nil)

    @options = utility.make_options.call(@_rootctx)

    if VoxgigStruct.getpath(@options, "feature.test.active") == true
      @mode = "test"
    end

    @_rootctx.options = @options

    # Add features from config.
    feature_opts = PoetrydbHelpers.to_map(VoxgigStruct.getprop(@options, "feature"))
    if feature_opts
      items = VoxgigStruct.items(feature_opts)
      if items
        items.each do |item|
          fname = item[0]
          fopts = PoetrydbHelpers.to_map(item[1])
          if fopts && fopts["active"] == true
            utility.feature_add.call(@_rootctx, PoetrydbFeatures.make_feature(fname))
          end
        end
      end
    end

    # Add extension features.
    extend_val = VoxgigStruct.getprop(@options, "extend")
    if extend_val.is_a?(Array)
      extend_val.each do |f|
        if f.respond_to?(:get_name)
          utility.feature_add.call(@_rootctx, f)
        end
      end
    end

    # Initialize features.
    @features.each do |f|
      utility.feature_init.call(@_rootctx, f)
    end

    utility.feature_hook.call(@_rootctx, "PostConstruct")
  end

  def options_map
    out = VoxgigStruct.clone(@options)
    out.is_a?(Hash) ? out : {}
  end

  def get_utility
    PoetrydbUtility.copy(@_utility)
  end

  def get_root_ctx
    @_rootctx
  end

  def prepare(fetchargs = {})
    utility = @_utility
    fetchargs ||= {}

    ctrl = PoetrydbHelpers.to_map(VoxgigStruct.getprop(fetchargs, "ctrl")) || {}

    ctx = utility.make_context.call({
      "opname" => "prepare",
      "ctrl" => ctrl,
    }, @_rootctx)

    opts = @options
    path = VoxgigStruct.getprop(fetchargs, "path") || ""
    path = "" unless path.is_a?(String)
    method_val = VoxgigStruct.getprop(fetchargs, "method") || "GET"
    method_val = "GET" unless method_val.is_a?(String)
    params = PoetrydbHelpers.to_map(VoxgigStruct.getprop(fetchargs, "params")) || {}
    query = PoetrydbHelpers.to_map(VoxgigStruct.getprop(fetchargs, "query")) || {}
    headers = utility.prepare_headers.call(ctx)

    base = VoxgigStruct.getprop(opts, "base") || ""
    base = "" unless base.is_a?(String)
    prefix = VoxgigStruct.getprop(opts, "prefix") || ""
    prefix = "" unless prefix.is_a?(String)
    suffix = VoxgigStruct.getprop(opts, "suffix") || ""
    suffix = "" unless suffix.is_a?(String)

    ctx.spec = PoetrydbSpec.new({
      "base" => base, "prefix" => prefix, "suffix" => suffix,
      "path" => path, "method" => method_val,
      "params" => params, "query" => query, "headers" => headers,
      "body" => VoxgigStruct.getprop(fetchargs, "body"),
      "step" => "start",
    })

    # Merge user-provided headers.
    uh = VoxgigStruct.getprop(fetchargs, "headers")
    if uh.is_a?(Hash)
      uh.each { |k, v| ctx.spec.headers[k] = v }
    end

    _, err = utility.prepare_auth.call(ctx)
    raise err if err

    utility.make_fetch_def.call(ctx)
  end

  def direct(fetchargs = {})
    utility = @_utility

    # direct() is the raw-HTTP escape hatch: it always returns a result hash
    # ({ "ok" => ..., ... }) and never raises. prepare() raises on error, so
    # trap that and surface it in the hash.
    begin
      fetchdef = prepare(fetchargs)
    rescue PoetrydbError => err
      return { "ok" => false, "err" => err }
    end

    fetchargs ||= {}
    ctrl = PoetrydbHelpers.to_map(VoxgigStruct.getprop(fetchargs, "ctrl")) || {}

    ctx = utility.make_context.call({
      "opname" => "direct",
      "ctrl" => ctrl,
    }, @_rootctx)

    url = fetchdef["url"] || ""
    fetched, fetch_err = utility.fetcher.call(ctx, url, fetchdef)

    return { "ok" => false, "err" => fetch_err } if fetch_err

    if fetched.nil?
      return {
        "ok" => false,
        "err" => ctx.make_error("direct_no_response", "response: undefined"),
      }
    end

    if fetched.is_a?(Hash)
      status = PoetrydbHelpers.to_int(VoxgigStruct.getprop(fetched, "status"))
      headers = VoxgigStruct.getprop(fetched, "headers") || {}

      # No-body responses (204, 304) and explicit zero content-length must
      # skip JSON parsing — calling json() on an empty body errors.
      content_length = headers.is_a?(Hash) ? headers["content-length"] : nil
      no_body = status == 204 || status == 304 || content_length.to_s == "0"

      json_data = nil
      unless no_body
        jf = VoxgigStruct.getprop(fetched, "json")
        if jf.is_a?(Proc)
          begin
            json_data = jf.call
          rescue StandardError
            # Non-JSON body — leave data nil, keep status/headers.
            json_data = nil
          end
        end
      end

      return {
        "ok" => status >= 200 && status < 300,
        "status" => status,
        "headers" => headers,
        "data" => json_data,
      }
    end

    return {
      "ok" => false,
      "err" => ctx.make_error("direct_invalid", "invalid response type"),
    }
  end


  # Idiomatic facade: client.author.list / client.author.load({ "id" => ... })
  def author
    require_relative 'entity/author_entity'
    @author ||= AuthorEntity.new(self, nil)
  end

  # Deprecated: use client.author instead.
  def Author(data = nil)
    require_relative 'entity/author_entity'
    AuthorEntity.new(self, data)
  end


  # Idiomatic facade: client.authorab.list / client.authorab.load({ "id" => ... })
  def authorab
    require_relative 'entity/authorab_entity'
    @authorab ||= AuthorabEntity.new(self, nil)
  end

  # Deprecated: use client.authorab instead.
  def Authorab(data = nil)
    require_relative 'entity/authorab_entity'
    AuthorabEntity.new(self, data)
  end


  # Idiomatic facade: client.combined_search.list / client.combined_search.load({ "id" => ... })
  def combined_search
    require_relative 'entity/combined_search_entity'
    @combined_search ||= CombinedSearchEntity.new(self, nil)
  end

  # Deprecated: use client.combined_search instead.
  def CombinedSearch(data = nil)
    require_relative 'entity/combined_search_entity'
    CombinedSearchEntity.new(self, data)
  end


  # Idiomatic facade: client.combined_search_with_field.list / client.combined_search_with_field.load({ "id" => ... })
  def combined_search_with_field
    require_relative 'entity/combined_search_with_field_entity'
    @combined_search_with_field ||= CombinedSearchWithFieldEntity.new(self, nil)
  end

  # Deprecated: use client.combined_search_with_field instead.
  def CombinedSearchWithField(data = nil)
    require_relative 'entity/combined_search_with_field_entity'
    CombinedSearchWithFieldEntity.new(self, data)
  end


  # Idiomatic facade: client.line.list / client.line.load({ "id" => ... })
  def line
    require_relative 'entity/line_entity'
    @line ||= LineEntity.new(self, nil)
  end

  # Deprecated: use client.line instead.
  def Line(data = nil)
    require_relative 'entity/line_entity'
    LineEntity.new(self, data)
  end


  # Idiomatic facade: client.linecount.list / client.linecount.load({ "id" => ... })
  def linecount
    require_relative 'entity/linecount_entity'
    @linecount ||= LinecountEntity.new(self, nil)
  end

  # Deprecated: use client.linecount instead.
  def Linecount(data = nil)
    require_relative 'entity/linecount_entity'
    LinecountEntity.new(self, data)
  end


  # Idiomatic facade: client.poemcount.list / client.poemcount.load({ "id" => ... })
  def poemcount
    require_relative 'entity/poemcount_entity'
    @poemcount ||= PoemcountEntity.new(self, nil)
  end

  # Deprecated: use client.poemcount instead.
  def Poemcount(data = nil)
    require_relative 'entity/poemcount_entity'
    PoemcountEntity.new(self, data)
  end


  # Idiomatic facade: client.random.list / client.random.load({ "id" => ... })
  def random
    require_relative 'entity/random_entity'
    @random ||= RandomEntity.new(self, nil)
  end

  # Deprecated: use client.random instead.
  def Random(data = nil)
    require_relative 'entity/random_entity'
    RandomEntity.new(self, data)
  end


  # Idiomatic facade: client.title.list / client.title.load({ "id" => ... })
  def title
    require_relative 'entity/title_entity'
    @title ||= TitleEntity.new(self, nil)
  end

  # Deprecated: use client.title instead.
  def Title(data = nil)
    require_relative 'entity/title_entity'
    TitleEntity.new(self, data)
  end


  # Idiomatic facade: client.titleab.list / client.titleab.load({ "id" => ... })
  def titleab
    require_relative 'entity/titleab_entity'
    @titleab ||= TitleabEntity.new(self, nil)
  end

  # Deprecated: use client.titleab instead.
  def Titleab(data = nil)
    require_relative 'entity/titleab_entity'
    TitleabEntity.new(self, data)
  end



  def self.test(testopts = nil, sdkopts = nil)
    sdkopts = sdkopts || {}
    sdkopts = VoxgigStruct.clone(sdkopts)
    sdkopts = {} unless sdkopts.is_a?(Hash)

    testopts = testopts || {}
    testopts = VoxgigStruct.clone(testopts)
    testopts = {} unless testopts.is_a?(Hash)
    testopts["active"] = true

    VoxgigStruct.setpath(sdkopts, "feature.test", testopts)

    sdk = PoetrydbSDK.new(sdkopts)
    sdk.mode = "test"
    sdk
  end
end
