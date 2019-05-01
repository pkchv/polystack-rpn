import decimal
import rpn

def compute(expression):
    try:
        solved = rpn.solve_rpn(expression)
        return str(solved.to_eng_string())
    except decimal.DivisionByZero:
        return "DivisionByZero"
    except decimal.InvalidOperation:
        return "InvalidOperation"
    except decimal.Overflow:
        return "Overflow"
    except:
        return "SyntaxError"
