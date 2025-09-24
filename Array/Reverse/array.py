
def reverse_arr(arr, start, end):
    if len(arr) == 0:
        return
    
    while start < end:
        arr[start], arr[end] = arr[end], arr[start]
        start+=1
        end -= 1

    print(arr)



if __name__ == "__main__":
    arr = [22,33,44,55,66]
    reverse_arr(arr, 0, len(arr) -1)
