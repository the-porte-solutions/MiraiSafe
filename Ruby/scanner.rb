require 'socket'
require 'timeout'

def is_port_open?(ip, port)
  begin
    Timeout::timeout(1) do
      begin
        s = TCPSocket.new(ip, port)
        s.close
        return true
      rescue Errno::ECONNREFUSED, Errno::EHOSTUNREACH
        return false
      end
    end
  rescue Timeout::Error
  end

  return false
end

puts "What is the IP? "
ip = gets
puts "Checking for port 23"
puts " . . . "
puts is_port_open?(ip, 22)
puts "Checking for port 2323"
puts " . . . "
puts is_port_open?(ip,2323)