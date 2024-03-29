#!/usr/bin/env ruby

require 'ffi'

module Strings
  extend FFI::Library
  ffi_lib File.expand_path("../strings.so", File.dirname(__FILE__))
  attach_function :Hello, :C_Hello, [:uint8], :string
end

puts "Running from Ruby"
puts Strings.Hello 0
puts Strings.Hello 1
