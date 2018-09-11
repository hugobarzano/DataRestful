#!/usr/bin/env python

def Calc(data,operator,val):
    result=[]
    value=float(val)
    for i in data:
        if operator == "+":
            result.append(float(i) + value)
        elif operator == "*":
            result.append(float(i) * value)
        elif operator == "-":
            result.append(float(i) - value)
        elif operator == "/":
            result.append(float(i) / value)
        else:
            print "Operator not suported"
    print "[SERVICE OPERATOR] CALC: "+ operator+ " result:"
    print result
    return result


