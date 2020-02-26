"""Main entrypint of the application."""
# Api factory import
import config
from flask import Flask
from matcha.views.user import *
from matcha.views.account import *

app = Flask(config.NAME)

@app.route("/", methods=['GET'])
def home():
    return "matcha"

@app.route("/user", methods=['GET'])
def get_users():
    return user.getUser(headers)

@app.route("/user/<uuid:uuidUser>", methods=['GET'])
def get_user(uuidUser):
    return user.getUser(headers, uuidUser=uuidUser)

@app.route("/account", methods=['GET'])
def get_accounts(uuidAccount):
    return account.getAccount(headers)

@app.route("/account/<uuid:uuidAccount>", methods=['GET'])
def get_account(uuidAccount):
    return account.getAccount(headers, uuidAccount=uuidAccount)

if __name__ == '__main__':
    app.run()
