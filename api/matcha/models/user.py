import psycopg2

def getUser(uuidUser=None):
    # check if JWT is admin here
    rqst = '''
        SELECT * FROM user
    '''
    if uuidUser is not None:
        rqst += 'WHERE Uuid = "{}"'.format(uuidUser)
    try:
        conn, cursor = get_conn_cursor()
        cursor.execute(rqst)
        u = cursor.fetchall() if uuidUser is None else cursor.fetchone()
        cursor.close()
        conn.close()
