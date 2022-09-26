package lib

import (
	"time"
)

const line_1 = "░░███╗░░██████╗░██████╗░░░██╗██╗███████╗░█████╗░███████╗░█████╗░░█████╗░░█████╗░██╗░░░"
const line_2 = "░████║░░╚════██╗╚════██╗░██╔╝██║██╔════╝██╔═══╝░╚════██║██╔══██╗██╔══██╗██╔══██╗╚═╝░░░" 
const line_3 = "██╔██║░░░░███╔═╝░█████╔╝██╔╝░██║██████╗░██████╗░░░░░██╔╝╚█████╔╝╚██████║██║░░██║░░░░░░" 
const line_4 = "╚═╝██║░░██╔══╝░░░╚═══██╗███████║╚════██╗██╔══██╗░░░██╔╝░██╔══██╗░╚═══██║██║░░██║░░░░░░" 
const line_5 = "███████╗███████╗██████╔╝╚════██║██████╔╝╚█████╔╝░░██╔╝░░╚█████╔╝░█████╔╝╚█████╔╝██╗░░░" 
const line_6 = "╚══════╝╚══════╝╚═════╝░░░░░░╚═╝╚═════╝░░╚════╝░░░╚═╝░░░░╚════╝░░╚════╝░░╚════╝░╚═╝░░░" 


var char_sequence = [...]string {line_1, line_2, line_3, line_4, line_5, line_6} 
const char_unit_height = 6 
const char_unit_width = 24
const char_semicollon_width = 9
const char_semicollon_pos = 11
const char_blank_pos = 12

func MapChars() map[int]string {
  var char_sequence_map = make(map[int]string)
  var number_char, semicollon_char, blank_char string

  for j := char_unit_width; j <= (char_unit_width * 10);  j = j + char_unit_width {
    for i := 0; i < char_unit_height; i++ {
      if i == 0 {
        number_char = char_sequence[i][j-char_unit_width:j]
      } else {
        number_char = number_char + char_sequence[i][j-char_unit_width:j] 
      }
    }
    if j/char_unit_width == 10 {
      char_sequence_map[0] = number_char
    }  else {
      char_sequence_map[j/char_unit_width] = number_char
    }
  }
  for i := 0; i < char_unit_height; i++ {
    if i == 0 {
      semicollon_char = char_sequence[i][(char_unit_width * 10):(char_unit_width * 10) + char_semicollon_width]
      blank_char      = char_sequence[i][(char_unit_width * 10) + char_semicollon_width :(char_unit_width * 10) + (char_semicollon_width * 2)]
    } else {
      semicollon_char = semicollon_char + char_sequence[i][(char_unit_width * 10):(char_unit_width * 10) + char_semicollon_width]
      blank_char = blank_char + char_sequence[i][(char_unit_width * 10) + char_semicollon_width :(char_unit_width * 10) + (char_semicollon_width * 2)]
    } 
  }
  char_sequence_map[char_semicollon_pos] = semicollon_char
  char_sequence_map[char_blank_pos] = blank_char
  return char_sequence_map
}

func PrintTime (t time.Time, tick bool) string {
  var hour_1, hour_2, min_1, min_2 int
	var index, index_sc int
	var result string
	var hour_string, minute_string, semicollon string

  my_sequence := MapChars()
	t_hour := t.Hour()
	t_min  := t.Minute()

	hour_1 = int(t_hour/10)
	hour_2 = t_hour - (hour_1 * 10)
	min_1 = int (t_min/10)
	min_2 = t_min - (min_1 * 10)

	for i := 1; i <= char_unit_height; i++ {
		index = (i * char_unit_width)
		index_sc =(i * char_semicollon_width)
		hour_string = my_sequence[hour_1][index - char_unit_width: index] +  
    my_sequence[hour_2][index - char_unit_width: index] 

		minute_string = my_sequence[min_1][index - char_unit_width: index] + 
		my_sequence[min_2][index - char_unit_width: index]
		if tick {
		  semicollon = my_sequence[char_semicollon_pos][index_sc  - char_semicollon_width: index_sc] 
		} else {
		  semicollon = my_sequence[char_blank_pos][index_sc  - char_semicollon_width: index_sc]
		}
		if i == char_unit_height {
			result = result +	hour_string + semicollon + minute_string 
		} else {
			result = result +	hour_string + semicollon + minute_string + "\n"
		}
	}
	// fmt.Println(result)

	return result
}
