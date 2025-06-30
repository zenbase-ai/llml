#!/usr/bin/env ruby

def bump_version(current_version, bump_type)
  parts = current_version.strip.split('.')

  unless parts.length == 3 && parts.all? { |p| p.match?(/^\d+$/) }
    raise "Invalid version format: #{current_version}. Expected X.Y.Z format."
  end

  major, minor, patch = parts.map(&:to_i)

  case bump_type.downcase
  when 'major'
    "#{major + 1}.0.0"
  when 'minor'
    "#{major}.#{minor + 1}.0"
  when 'patch'
    "#{major}.#{minor}.#{patch + 1}"
  when /^\d+\.\d+\.\d+$/
    bump_type
  else
    raise "Invalid bump type: #{bump_type}. Use 'major', 'minor', 'patch', or 'X.Y.Z'."
  end
end

if ARGV.length != 1
  puts "Usage: #{$0} <major|minor|patch|X.Y.Z>"
  exit 1
end

current_version = STDIN.read
new_version = bump_version(current_version, ARGV[0])
puts new_version
