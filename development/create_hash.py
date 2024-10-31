# パスワードをハッシュ化するための関数
# DBに仮のパスワードを挿入するときに使用する
# bcryptはインストールする必要あり

import bcrypt

def hash_password(password):
    password_bytes = password.encode('utf-8')
    hashed = bcrypt.hashpw(password_bytes, bcrypt.gensalt())
    return hashed

# 使用例
password = "yourpassword" #ここにパスワードを入れる
hashed_password = hash_password(password)

# ハッシュ化されたパスワードを表示
print(hashed_password.decode('utf-8'))  # デコードして文字列として表示