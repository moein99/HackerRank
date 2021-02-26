package main

import (
    "fmt"
    "bufio"  
    "os"
    "io"
    "strings"
    "strconv"
)

func sumStringSlice(nums []string) (int, error) {
    total := 0
    
    for i := 0; i < len(nums); i++ {
        num, err := strconv.Atoi(nums[i])
        if err != nil {
            return -1, err
        }
        total += num
    }
    
    return total, nil
}

func readOneLineOfNums() ([]string, error) {
    reader := bufio.NewReader(os.Stdin)
    nums_str, err := reader.ReadString('\n')
    if err != io.EOF {
        return nil, err;
    }
    return strings.Split(nums_str, " "), nil
}

func main() {
    var i int;
    _, err := fmt.Scanf("%d", &i)
    if err != nil {
        return;
    }
    
    nums, err := readOneLineOfNums()
    if err != nil {
        return;
    }
    
    total, err := sumStringSlice(nums)
    if err != nil {
        return;
    }
    
    fmt.Print(total)
}

