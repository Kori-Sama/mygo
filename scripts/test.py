import requests

BASE_URL = 'http://localhost:8888/api'

headers

def register():
    register_url = BASE_URL + '/register'
    register_data = {
        'username': 'testuser',
        'password': '123',
        'role': 'Admin'
    }
    response = requests.post(register_url, json=register_data)

    print(response.body)

    if response.status_code == 200:
        print("Registration Successful")
    else:
        print("Registration Failed")

    token = response.headers.get('Authorization')

    headers = {'Authorization': token}

def getSelfTransactions():
    get_self_transactions_url = BASE_URL + '/transaction/self'
    response = requests.get(get_self_transactions_url, headers=headers)

    print(response.body)

    if response.status_code == 200:
        print("Get Self Transactions Successful")
    else:
        print("Get Self Transactions Failed")

if __name__ == "__main__":
    register()
    #delay 1000ms
    
    getSelfTransactions()