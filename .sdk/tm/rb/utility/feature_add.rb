# Poetrydb SDK utility: feature_add
module PoetrydbUtilities
  FeatureAdd = ->(ctx, f) {
    ctx.client.features << f
  }
end
