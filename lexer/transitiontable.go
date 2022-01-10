
package lexer



/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates] func(rune) int

var TransTab = TransitionTable{
	
		// S0
		func(r rune) int {
			switch {
			case r == 9 : // ['\t','\t']
				return 1
			case r == 10 : // ['\n','\n']
				return 1
			case r == 13 : // ['\r','\r']
				return 1
			case r == 32 : // [' ',' ']
				return 1
			case r == 33 : // ['!','!']
				return 2
			case r == 34 : // ['"','"']
				return 3
			case r == 38 : // ['&','&']
				return 4
			case r == 40 : // ['(','(']
				return 5
			case r == 41 : // [')',')']
				return 6
			case r == 42 : // ['*','*']
				return 7
			case r == 43 : // ['+','+']
				return 8
			case r == 45 : // ['-','-']
				return 9
			case r == 47 : // ['/','/']
				return 10
			case 48 <= r && r <= 57 : // ['0','9']
				return 11
			case r == 58 : // [':',':']
				return 12
			case r == 60 : // ['<','<']
				return 13
			case r == 61 : // ['=','=']
				return 14
			case r == 62 : // ['>','>']
				return 15
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 97 : // ['a','a']
				return 16
			case r == 98 : // ['b','b']
				return 17
			case 99 <= r && r <= 100 : // ['c','d']
				return 16
			case r == 101 : // ['e','e']
				return 18
			case r == 102 : // ['f','f']
				return 19
			case 103 <= r && r <= 104 : // ['g','h']
				return 16
			case r == 105 : // ['i','i']
				return 20
			case 106 <= r && r <= 114 : // ['j','r']
				return 16
			case r == 115 : // ['s','s']
				return 21
			case r == 116 : // ['t','t']
				return 22
			case 117 <= r && r <= 122 : // ['u','z']
				return 16
			case r == 123 : // ['{','{']
				return 23
			case r == 124 : // ['|','|']
				return 24
			case r == 125 : // ['}','}']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S1
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S2
		func(r rune) int {
			switch {
			case r == 61 : // ['=','=']
				return 26
			
			
			
			}
			return NoState
			
		},
	
		// S3
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 27
			
			
			default:
				return 3
			}
			
		},
	
		// S4
		func(r rune) int {
			switch {
			case r == 38 : // ['&','&']
				return 28
			
			
			
			}
			return NoState
			
		},
	
		// S5
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S6
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S7
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S8
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S9
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S10
		func(r rune) int {
			switch {
			case r == 42 : // ['*','*']
				return 29
			case r == 47 : // ['/','/']
				return 30
			
			
			
			}
			return NoState
			
		},
	
		// S11
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S12
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S13
		func(r rune) int {
			switch {
			case r == 61 : // ['=','=']
				return 31
			
			
			
			}
			return NoState
			
		},
	
		// S14
		func(r rune) int {
			switch {
			case r == 61 : // ['=','=']
				return 32
			
			
			
			}
			return NoState
			
		},
	
		// S15
		func(r rune) int {
			switch {
			case r == 61 : // ['=','=']
				return 33
			
			
			
			}
			return NoState
			
		},
	
		// S16
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S17
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 110 : // ['a','n']
				return 16
			case r == 111 : // ['o','o']
				return 34
			case 112 <= r && r <= 122 : // ['p','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S18
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 107 : // ['a','k']
				return 16
			case r == 108 : // ['l','l']
				return 35
			case 109 <= r && r <= 122 : // ['m','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S19
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case r == 97 : // ['a','a']
				return 36
			case 98 <= r && r <= 110 : // ['b','n']
				return 16
			case r == 111 : // ['o','o']
				return 37
			case 112 <= r && r <= 122 : // ['p','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S20
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 101 : // ['a','e']
				return 16
			case r == 102 : // ['f','f']
				return 38
			case 103 <= r && r <= 109 : // ['g','m']
				return 16
			case r == 110 : // ['n','n']
				return 39
			case 111 <= r && r <= 122 : // ['o','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S21
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 115 : // ['a','s']
				return 16
			case r == 116 : // ['t','t']
				return 40
			case 117 <= r && r <= 122 : // ['u','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S22
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 113 : // ['a','q']
				return 16
			case r == 114 : // ['r','r']
				return 41
			case 115 <= r && r <= 122 : // ['s','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S23
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S24
		func(r rune) int {
			switch {
			case r == 124 : // ['|','|']
				return 42
			
			
			
			}
			return NoState
			
		},
	
		// S25
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S26
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S27
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S28
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S29
		func(r rune) int {
			switch {
			case r == 42 : // ['*','*']
				return 43
			
			
			default:
				return 29
			}
			
		},
	
		// S30
		func(r rune) int {
			switch {
			case r == 10 : // ['\n','\n']
				return 44
			
			
			default:
				return 30
			}
			
		},
	
		// S31
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S32
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S33
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S34
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 110 : // ['a','n']
				return 16
			case r == 111 : // ['o','o']
				return 45
			case 112 <= r && r <= 122 : // ['p','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S35
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 114 : // ['a','r']
				return 16
			case r == 115 : // ['s','s']
				return 46
			case 116 <= r && r <= 122 : // ['t','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S36
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 107 : // ['a','k']
				return 16
			case r == 108 : // ['l','l']
				return 47
			case 109 <= r && r <= 122 : // ['m','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S37
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 113 : // ['a','q']
				return 16
			case r == 114 : // ['r','r']
				return 48
			case 115 <= r && r <= 122 : // ['s','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S38
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S39
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 115 : // ['a','s']
				return 16
			case r == 116 : // ['t','t']
				return 49
			case 117 <= r && r <= 122 : // ['u','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S40
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 113 : // ['a','q']
				return 16
			case r == 114 : // ['r','r']
				return 50
			case 115 <= r && r <= 122 : // ['s','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S41
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 116 : // ['a','t']
				return 16
			case r == 117 : // ['u','u']
				return 51
			case 118 <= r && r <= 122 : // ['v','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S42
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S43
		func(r rune) int {
			switch {
			case r == 42 : // ['*','*']
				return 43
			case r == 47 : // ['/','/']
				return 52
			
			
			default:
				return 29
			}
			
		},
	
		// S44
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S45
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 107 : // ['a','k']
				return 16
			case r == 108 : // ['l','l']
				return 53
			case 109 <= r && r <= 122 : // ['m','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S46
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 100 : // ['a','d']
				return 16
			case r == 101 : // ['e','e']
				return 54
			case 102 <= r && r <= 122 : // ['f','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S47
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 114 : // ['a','r']
				return 16
			case r == 115 : // ['s','s']
				return 55
			case 116 <= r && r <= 122 : // ['t','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S48
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 108 : // ['a','l']
				return 16
			case r == 109 : // ['m','m']
				return 56
			case 110 <= r && r <= 122 : // ['n','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S49
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 100 : // ['a','d']
				return 16
			case r == 101 : // ['e','e']
				return 57
			case 102 <= r && r <= 122 : // ['f','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S50
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 104 : // ['a','h']
				return 16
			case r == 105 : // ['i','i']
				return 58
			case 106 <= r && r <= 122 : // ['j','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S51
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 100 : // ['a','d']
				return 16
			case r == 101 : // ['e','e']
				return 59
			case 102 <= r && r <= 122 : // ['f','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S52
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S53
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 100 : // ['a','d']
				return 16
			case r == 101 : // ['e','e']
				return 60
			case 102 <= r && r <= 122 : // ['f','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S54
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S55
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 100 : // ['a','d']
				return 16
			case r == 101 : // ['e','e']
				return 61
			case 102 <= r && r <= 122 : // ['f','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S56
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S57
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 102 : // ['a','f']
				return 16
			case r == 103 : // ['g','g']
				return 62
			case 104 <= r && r <= 122 : // ['h','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S58
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 109 : // ['a','m']
				return 16
			case r == 110 : // ['n','n']
				return 63
			case 111 <= r && r <= 122 : // ['o','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S59
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S60
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case r == 97 : // ['a','a']
				return 64
			case 98 <= r && r <= 122 : // ['b','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S61
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S62
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 100 : // ['a','d']
				return 16
			case r == 101 : // ['e','e']
				return 65
			case 102 <= r && r <= 122 : // ['f','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S63
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 102 : // ['a','f']
				return 16
			case r == 103 : // ['g','g']
				return 66
			case 104 <= r && r <= 122 : // ['h','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S64
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 109 : // ['a','m']
				return 16
			case r == 110 : // ['n','n']
				return 67
			case 111 <= r && r <= 122 : // ['o','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S65
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 113 : // ['a','q']
				return 16
			case r == 114 : // ['r','r']
				return 68
			case 115 <= r && r <= 122 : // ['s','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S66
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S67
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
		// S68
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 16
			case 65 <= r && r <= 90 : // ['A','Z']
				return 16
			case r == 95 : // ['_','_']
				return 16
			case 97 <= r && r <= 122 : // ['a','z']
				return 16
			
			
			
			}
			return NoState
			
		},
	
}
