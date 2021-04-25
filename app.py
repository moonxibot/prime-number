def factor(num):
    result = []
    for i in range(1, num+1):
        if num % i == 0:
            result.append(i)

    return result

def prime(num):
    result = len(factor(num))
    if result == 2:
        return True
    elif result > 2:
        return False
    elif result < 2:
        return None

num = 2419
data = factor(num)
value = []

for i in data:
    if prime(i):
        value.append(i)

print(data)
print(value)
