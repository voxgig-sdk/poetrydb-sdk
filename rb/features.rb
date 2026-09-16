# Poetrydb SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/ratelimit_feature'
require_relative 'feature/retry_feature'
require_relative 'feature/test_feature'
require_relative 'feature/timeout_feature'


module PoetrydbFeatures
  def self.make_feature(name)
    case name
    when "base"
      PoetrydbBaseFeature.new
    when "ratelimit"
      PoetrydbRatelimitFeature.new
    when "retry"
      PoetrydbRetryFeature.new
    when "test"
      PoetrydbTestFeature.new
    when "timeout"
      PoetrydbTimeoutFeature.new
    else
      PoetrydbBaseFeature.new
    end
  end
end
