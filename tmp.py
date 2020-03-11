import csv
import sys


with open('user_id_gender.csv') as f:
    try:
        reader = csv.reader(f, delimiter=',', quotechar='"')
    except Exception as e:
        sys.exit(1)
    s = ""
    base = "INSERT INTO images(users_id, image) VALUES "
    image_id_female = 0
    image_id_male = 0
    for r in reader:
        if r[1] == 'F':
            path = '{}_women.jpg'.format(image_id_female)
            image_id_female = image_id_female +1 if image_id_female < 94 else 0
        else:
            path = '{}_man.jpg'.format(image_id_male)
            image_id_male = image_id_male + 1 if image_id_male < 94 else 0
        s += base + "({},bytea('{}'));\n".format(r[0], path)
    print(s)
