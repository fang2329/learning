function GetStr()
        print "world"
end

function GetBigger(a,b)
	if a >= b then
		print (a)
	else
		print (b)
	end
end


function GetResult()
	return "hello"
end

function Compare(a,b)
	if a >= b then
		return a
	else
		return b
	end
end

function MoreReturn(a)
	if (a == 10) then
		return "world","hello","golang"
	end
end
