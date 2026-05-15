# Poetrydb SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module PoetrydbFeatures
  def self.make_feature(name)
    case name
    when "base"
      PoetrydbBaseFeature.new
    when "test"
      PoetrydbTestFeature.new
    else
      PoetrydbBaseFeature.new
    end
  end
end
