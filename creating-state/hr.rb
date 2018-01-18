def permutate(string1, string2)
    return 0 if string1.length == 0 || string2.length == 0 #|| string1.length != string2.length
    string1_histogram = build_histogram(string1)

    compare_string_and_histogram(string2, string1_histogram)
end

def compare_string_and_histogram(string, historgram)
    string.chars.each do |c|
        if histogram.has_key?(c)
            histogram[c] -= 1
            return 0 if histogram[c] < 0
        else
            return 0
        end
    end

    1
end

def build_histogram(string)
    histogram = {}

    string.chars.each do |c|
        if histogram.has_key?(c)
            histogram[c] += 1
        else
            histogram[c] = 1
        end
    end

    histogram
end

def carParking(n, available)  
    result = [101,101]
    available.each.with_index(1) do |s, i|
        num_taken = s.count(1)
        first_open = s.index(0)
        return [i, 1] if num_taken == 0
       
        if num_taken < result[1] && first_open
            result[0] = i
            result[1] = first_open
        end
    end
    
    puts result
end

arr = [
    [1,1,1,1,1],
    [1,1,1,1,1],
    [1,1,1,1,1],
    [1,1,1,1,1],
    [1,1,1,0,0]
]

arr = [[0]
carParking(5, arr)


