import requests as re

resp = re.post("https://web.njit.edu/~gmo9/back-end/backend.php", data={"username": "jane","password": "apple"})
print(resp.text)