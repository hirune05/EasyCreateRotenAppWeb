
# 説明
# バックエンドへのアクセスに必要なJWTトークンを作成するpythonスクリプト

# 使用方法
# 0. backend/.envファイルに JWT_SECRET_KEY を設定する
# 1. カレントディレクトリを development にする。 もしくは、31行目のenv_var = load_env()の括弧内に.envファイルのパスを書く
# 3. コマンド入力(--id int はuserIdの指定, 省略可能)     python create_token.py --id 112233
# 2. 出力された以下のような文字列をcurlのヘッダに追加する
# Cookie: Authorization=XXXXXXXXXXXJWT_TOKENXXXXXXXXXXXXx


import base64
import json
import hmac
import hashlib
import time
import configparser
import argparse

# .envファイルを読み取る関数
def load_env(filename="../backend/.env"):
    with open(filename, "r") as f:
      config_data = f.readlines()
    config_data = list(map(lambda x: x[:-1] if x[-1] == "\n" else x,filter(lambda x: len(x) > 14 and x[0:14] == "JWT_SECRET_KEY", config_data)))
    if len(config_data) != 1:
      raise ValueError(f"'{filename}' file not include only JWT_SECRET_KEY")
    config_data = config_data[0].replace(" ", "").split("=")
    return config_data[1]

# 環境変数を読み取る
env_var = load_env()
SECRET_KEY = env_var.encode()  # バイト列に変換

# コマンドライン引数を解析する関数
def parse_arguments():
    parser = argparse.ArgumentParser(description="Generate JWT token.")
    parser.add_argument('--id', type=int, default=123, help="User ID")
    return parser.parse_args()

# コマンドライン引数を取得
args = parse_arguments()
user_id = args.id

# ペイロード（トークンに含めるデータ）
payload = {
    "id": user_id,  # ユーザーIDをコマンドライン引数から取得
    "exp": int(time.time()) + 3600  # 有効期限（1時間後）
}

# JWTのヘッダーを作成
header = {
    "alg": "HS256",
    "typ": "JWT"
}

# Base64URLエンコーディングの関数
def base64url_encode(data):
    json_data = json.dumps(data, separators=(',', ':')).encode()
    return base64.urlsafe_b64encode(json_data).rstrip(b'=')

# ヘッダーとペイロードをエンコード
encoded_header = base64url_encode(header)
encoded_payload = base64url_encode(payload)

# 署名を作成
signature_input = f"{encoded_header.decode()}.{encoded_payload.decode()}"
signature = hmac.new(SECRET_KEY, signature_input.encode(), hashlib.sha256).digest()
encoded_signature = base64.urlsafe_b64encode(signature).rstrip(b'=')

# JWTトークンを生成
jwt_token = f"{encoded_header.decode()}.{encoded_payload.decode()}.{encoded_signature.decode()}"

print(f"Cookie: Authorization={jwt_token}")

