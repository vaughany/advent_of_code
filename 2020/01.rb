#!/usr/bin/env ruby

# Advent of Code 2020. Day One. http://adventofcode.com/2020/day/1

require 'benchmark'

if RUBY_VERSION != '2.7.0' then
  puts "Error, quitting: please use Ruby 2.7: `rvm use ruby-2.7.0`."
  exit(1)
end

def foo
  time = Benchmark.measure {
    puts "Using: Ruby #{RUBY_VERSION}."
  }

  puts time # puts time.real
end

foo
