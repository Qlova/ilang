package main

import (
	"fmt"
	"text/scanner"
	"io"
	"os"
	//"strconv"
	"strings"
)

//This is an expression shunter. It takes the current identifyer and shunts it into the next operator.
//I don't think this has anything to do with the shunting yard algorithim, I just like the term.
func shunt(name string, s *scanner.Scanner, output io.Writer) string {

		//Scan the next token.
		s.Scan()
		
		//If it is one of these characters, then we have finished our shunt.
		switch s.TokenText() {
			case ")", ",", "\n", "]", ";":
				return name
		}
		
		var token = s.TokenText()
		
		//I love doing the shunting.
		if IsOperator(token+string(s.Peek())) {
			token += string(s.Peek())
			s.Scan()
		}
		
		//What a shunting mess.
		if IsOperator(token) {
			s.Scan()
			//fmt.Println(s.TokenText())
			
			unique++
			id := "i+shunt+"+fmt.Sprint(unique)
			
			var operator Operator
			var ok bool
			
			var A = ExpressionType
			var B TYPE
			var next string
			if operator, ok = GetOperator(token, ExpressionType, UNDEFINED); !ok {
			
				next = expression(s, output, OperatorPrecident(token))
				B = ExpressionType
				operator, ok = GetOperator(token, A, B)
				if token == "=" && A == STRING && B == STRING {
					ExpressionType = NUMBER
				}
			} else if token == "²" {
				next = name
				B = ExpressionType
				operator, ok = GetOperator("*", A, B)
				token = "*"
			}

			if ok {
				
				asm := operator.Assembly
				asm = strings.Replace(asm, "%a", name, -1)
				asm = strings.Replace(asm, "%b", next, -1)
				asm = strings.Replace(asm, "%c", id, -1)
				
				fmt.Fprint(output, asm, "\n")
				
				if !OperatorPrecident(token) {
					return shunt(id, s, output)
				}
				return id
				
			} else {
				fmt.Println(s.Pos(), "Invalid Operator Matchup! ", A , token, B, "(types do not support the opperator)")
				os.Exit(1)
			}
		}
		
		if s.TokenText() == "." {
			s.Scan()
			return shunt(IndexUserType(s, output, name, s.TokenText()), s, output)
		}
		
		if s.TokenText() == "[" {
			s.Scan()
			if ExpressionType != STRING && ExpressionType < USER {
				RaiseError(s, "Cannot index "+name+", not an array! ("+ExpressionType.String()+")")
			}	
			
			var index = expression(s, output)
			
			ExpressionType = NUMBER
			
			unique++
			
			fmt.Fprintf(output, "PLACE %v\nPUSH %v\nGET %v\n", name, index, "i+shunt+"+fmt.Sprint(unique))
			return shunt("i+shunt+"+fmt.Sprint(unique), s, output)
		}
		
		fmt.Println(s.Pos(), "[SHUNTING YARD] Unexpected ", s.TokenText(), "("+name+")")
		os.Exit(1)
		return ""
}
