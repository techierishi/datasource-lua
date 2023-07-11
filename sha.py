from hashlib import md5
m = md5()
with open("techierishi-luaquery-datasource-1.0.0.zip", "rb") as f:
    data = f.read()
    m.update(data)
    print(m.hexdigest())