import os

NAME = "Matcha"
TITLE = "Matcha Backend API."
VERSION = "0.1.0"
DESCRIPTION = "Matcha Backend API."

DEBUG = False
TESTING = True
CSRF_ENABLED = True
SECRET_KEY_JWT = os.environ.get('SECRET_KEY_JWT', 'secret')
DB_USER = os.environ.get('DB_USER', 'admin')
DB_PASSWORD = os.environ.get('DB_PASSWORD', 'secret')
DB_HOST = os.environ.get('DB_HOST', 'db-matcha')
DB_DBNAME = os.environ.get('DB_DBNAME', 'matcha')
DB_PORT = os.environ.get('DB_PORT', 5432)
