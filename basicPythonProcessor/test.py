#!/usr/bin/env python

from operatorBasic import Calc




def TestBasicSum():
    data = ["1.0", "2.0", "3.0"]
    expected=[2.0,3.0,4.0]
    operator = "+"
    value = "1"
    testcase=Calc(data, operator, value)
    if testcase[0]== expected[0] and testcase[1]== expected[1] and testcase[2]== expected[2]:
        print "[SERVICE OPERATOR TEST] Sum Pass"
        return True
    else:
        print "[SERVICE OPERATOR TEST] Sum Fail"
        return False



def TestBasicSub():
    data = ["1.0", "2.0", "3.0"]
    expected=[0.0,1.0,2.0]
    operator = "-"
    value = "1"
    testcase=Calc(data, operator, value)
    if testcase[0]== expected[0] and testcase[1]== expected[1] and testcase[2]== expected[2]:
        print "[SERVICE OPERATOR TEST] Sub Pass"
        return True
    else:
        print "[SERVICE OPERATOR TEST] Sub Fail"
        return False



def TestBasicMul():
    data = ["1.0", "2.0", "3.0"]
    expected=[2.0,4.0,6.0]
    operator = "*"
    value = "2"
    testcase=Calc(data, operator, value)
    if testcase[0]== expected[0] and testcase[1]== expected[1] and testcase[2]== expected[2]:
        print "[SERVICE OPERATOR TEST] Mul Pass"
        return True
    else:
        print "[SERVICE OPERATOR TEST] Mul Fail"
        return False




def TestBasicDiv():
    data = ["2.0", "4.0", "8.0"]
    expected=[1.0,2.0,4.0]
    operator = "/"
    value = "2"
    testcase=Calc(data, operator, value)
    if testcase[0]== expected[0] and testcase[1]== expected[1] and testcase[2]== expected[2]:
        print "[SERVICE OPERATOR TEST] Div Pass"
        return True
    else:
        print "[SERVICE OPERATOR TEST] Div Fail"
        return False






TestBasicSum()
TestBasicSub()
TestBasicMul()
TestBasicDiv()



