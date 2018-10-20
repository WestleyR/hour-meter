// Created by: Westley K
// email: westley@sylabs.io
// Date: Oct 20, 2018
// https://github.com/WestleyK/hour-meter
// Version-1.0.0
//
// MIT License
//
// Copyright (c) 2018 WestleyK
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//


package main


import (
    "fmt"
    "os"
    "time"
    "strings"
    "strconv"
    "io/ioutil"
)


var (
    SCRIPT_VERSION string = "version-1.0.0"
    SCRIPT_DATE string = "Oct 20, 2018"

    SCRIPT_NAME string = ""
    OPTION string = ""

    START_TIME int64 = 0
    WAIT_TIME int64 = 0

    TIME_FILE string = ".time-meter.txt"
    FILE string = ""
    TIME string = ""
    MINUT int64 = 0
    HOUR int64 = 0
    minute_left int64 = 0

    h_set bool = false
    m_set bool = false
    set_minute bool = false

    yes_output bool = false
    NEW_FILE bool = false

)

func print_version() {
    fmt.Println(SCRIPT_VERSION)
    fmt.Println(SCRIPT_DATE)
    os.Exit(0)
}

func print_help() {
    fmt.Print("Usage: ", SCRIPT_NAME, " [OPTION] [FILE]\n")
    fmt.Print("    -h, -help, --help, help (print help menu)\n")
    fmt.Print("    -i, -info, --info, info (print script info)\n")
    fmt.Print("    -o (print hour : minute every minute)\n")
    fmt.Print("    -file [FILE] (show default file, or use specified file)\n")
    fmt.Print("    -v, -version, --version, version (print script version)\n")
    fmt.Print("source code: https://github.com/WestleyK/hour-meter\n")
    os.Exit(0)
}

func print_info() {
    info()
    os.Exit(0)
}

func fail(bug string, err string) {
    fmt.Print(bug)
    if len(err) >= 1 {
        fmt.Print(err)
    }
    fmt.Print("\n")
    os.Exit(1)
}

func start() {
    b, err := ioutil.ReadFile(TIME_FILE)
    if err != nil {
        fmt.Print(err)
    }

    str := string(b)

    result := strings.Split(str, "  ")
    
    for i := range result {
        if _, err := strconv.Atoi(result[i]); err == nil {
            if h_set != true {
                HOUR, _ = strconv.ParseInt(result[i], 10, 64)
//                fmt.Println(HOUR)
                //HOUR, err = strconv.Atoi(result[i])
                h_set = true
            } else if m_set != true {
                MINUT, _ = strconv.ParseInt(result[i], 10, 64)
                minute_left, _ = strconv.ParseInt(result[i], 10, 64)
                set_minute = true
                //MINUT, err = strconv.Atoi(result[i])
                m_set = true
            }
        }
    }

//    fmt.Println(HOUR)
//    fmt.Println(MINUT)

}

func write_file(TIME string) {
    file, err := os.Create(TIME_FILE)
        if err != nil {
            fmt.Print("ERROR: ")
            fmt.Println(err)
            os.Exit(1)
        }
    defer file.Close()
    fmt.Fprintf(file, TIME)
}

func option_file() {
    if len(os.Args) >= 3 {
    //if len(os.Args[3:]) >= 1 {
        NEW_FILE = true
        TIME_FILE = os.Args[2]
        fmt.Print("new file: ", TIME_FILE, "\n")
        return
    }
    fmt.Println(TIME_FILE)
    os.Exit(0)
}

func check_args(OPTION string) {
    if OPTION == "-h" || OPTION == "-help" || OPTION == "--help" || OPTION == "help" {
        print_help()
        return
    } else if OPTION == "-i" || OPTION == "-info" || OPTION == "--info" || OPTION == "info" {
        print_info()
        return
    } else if OPTION == "-o" || OPTION == "-out" {
        yes_output = true
        return
    } else if OPTION == "-file" {
        option_file()
        return
    } else if OPTION == "-v" || OPTION == "-version" || OPTION == "--version" || OPTION == "version" {
        print_version()
        return
    } else {
        fail("option not found! ", OPTION)
    }
}

func main() {
    START_TIME = time.Now().Unix()
    SCRIPT_NAME = os.Args[0]

    if len(os.Args[1:]) >= 1 {
        check_args(os.Args[1])
    }

    if len(os.Args) >= 4 {
        fail("too many arguments!", "")
    }

    start()
    TIME_START := time.Now().Unix()
    for {
        time.Sleep(1 * time.Second)

        if set_minute == true {
            MINUT = time.Now().Unix() - TIME_START
            MINUT += minute_left
        } else {
            MINUT = time.Now().Unix() - TIME_START
        }


        if MINUT >= 6 {
            HOUR += 1
            MINUT = 0
            TIME_START = time.Now().Unix()
            set_minute = false
        }

    
//        hour_string := strconv.Itoa(HOUR)
        hour_string := strconv.FormatInt(HOUR, 10)
//        minut_string := strconv.Itoa(MINUT)
        minut_string := strconv.FormatInt(MINUT, 10)
        time_string := []string{hour_string, " hours and ", minut_string, " minutes\n"}
        TIME = (strings.Join(time_string, " "))
        if yes_output == true {
            fmt.Print(TIME)
        }
        //fmt.Print(TIME)
        write_file(TIME)

    }


}





//
// End source code
//

