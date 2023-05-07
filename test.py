import json
import random

import requests


def generate_random_str(base_str="",randomlength=16):
    """
    生成一个指定长度的随机字符串
    """
    random_str = ''
    length = len(base_str) - 1
    for i in range(randomlength):
        random_str += base_str[random.randint(0, length)]
    return random_str


def username_list():
    list01 = []
    base_str = 'ABCDEFGHIGKLMNOPQRSTUVWXYZabcdefghigklmnopqrstuvwxyz'
    for i in range(10):
        s = generate_random_str(base_str,6)
        list01.append(s)
    return list01


def password_list():
    list01 = []
    for i in range(10):
        list01.append("123456")
    return list01
def Phone_list():
    list01 = []
    base_str = '0123456789'
    for i in range(10):
        s=generate_random_str(base_str,11)
        list01.append(s)
    return list01
def AgeList():
    list01 = []
    for i in range(15,25):
        list01.append(i)
    return list01
def sex_list():
    list=["男","女"]

    list01=[list[random.randint(0,1)] for i in range(10)]
    return  list01

if __name__ == '__main__':
    usernameList=username_list()
    passwordList=password_list()
    phonelist=Phone_list()
    agelist=AgeList()
    sexlist=sex_list()
    print(usernameList)
    print(passwordList)
    print(phonelist)
    print(agelist)
    print(sexlist)
    headers = {
        'User-Agent': 'Apipost client Runtime/+https://www.apipost.cn/',
        'Content-Type': 'application/json',
    }

    for i in  range(10):
        data={
            "username":usernameList[i],
            "password":passwordList[i],
            "sex":sexlist[i],
            "age":agelist[i],
            "phone":phonelist[i],
        }
        # print(data)
        a=requests.post('http://127.0.0.1:8080/register', headers=headers, data=json.dumps(data))
        # print(a.text)
