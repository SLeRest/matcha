
def getUser(uuidUser=None):
    # check if JWT is admin here
    rqst = '''
        SELECT * FROM user
    '''
    if uuidUser is not None:
        rqst += 'WHERE Uuid = "{}"'.format(uuidUser)
    conn, cursor = get_conn_cursor()
    cursor.execute(rqst)
    conn.commit(rqst)
    cursor.close()
    conn.close()
