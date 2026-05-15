# Poetrydb SDK utility: make_context
require_relative '../core/context'
module PoetrydbUtilities
  MakeContext = ->(ctxmap, basectx) {
    PoetrydbContext.new(ctxmap, basectx)
  }
end
