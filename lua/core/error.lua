-- Poetrydb SDK error

local PoetrydbError = {}
PoetrydbError.__index = PoetrydbError


function PoetrydbError.new(code, msg, ctx)
  local self = setmetatable({}, PoetrydbError)
  self.is_sdk_error = true
  self.sdk = "Poetrydb"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function PoetrydbError:error()
  return self.msg
end


function PoetrydbError:__tostring()
  return self.msg
end


return PoetrydbError
