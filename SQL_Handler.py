import mysql.connector
import random
import re


class Handle_data:
    def __init__(self, seed):
        # Creating connection object
        self.conn = mysql.connector.connect(user="Persian_dvl", db='Persian', password='password', auth_plugin='mysql_native_password')
        self.cursor = self.conn.cursor()
        random.seed(seed)
        self.image_path = "/home/charlotte/Bureau/Projects/Persian/images/"

    def date_format(self, date):
        """Changes the date format between external and internal format"""
        # Internal format to external format
        if '-' in date:
            date = str(date[8:10] + '/' + date[5:7] + '/' + date[0:4])
        # External format to internal format
        else:
            date = str(date[6:10] + '-' + date[3:5] + '-' + date[0:2])
        return date

    def add_lecture(self, nr, date):
        """Adds a lecture to the 'Lecture' table.
        INPUT:  nr: number of the lecture
                date: date of the lecture in external format"""
        date = self.date_format(date)
        # Create a new lecture
        self.cursor.execute("INSERT INTO Lecture VALUES ('{}', '{}')".format(nr, date))
        self.conn.commit()

    def add_word(self, lst):
        """Adds a word to the 'Words' table
        INPUT: lst: list of the word (in french), the phonetic, the persian translation"""
        self.cursor.execute("SELECT MAX(woid) FROM Words")
        last_id = self.cursor.fetchall()
        if last_id[0][0] == None:
            new_id = 1
        else:
            new_id = last_id[0][0] + 1
        new_lst = [new_id, lst[0], lst[1], lst[2], lst[3]]
        # Create a new word
        self.cursor.execute("INSERT INTO Words VALUES ('{}', '{}', '{}', '{}', '{}', 'N')".format(new_lst[0], new_lst[1], new_lst[2], new_lst[3], new_lst[4]))
        self.conn.commit()

    def import_words(self, file):
        leid = re.sub(r'^.*?_', '', file)
        leid = leid[:-4]
        with open(file, 'r') as f:
            lines = f.readlines()
            for i in lines:
                i = i.rstrip()
                i = i.split(",")
                self.add_word([i[0], i[1], i[2], leid])


    def get_words_lecture(self, leid):
        """Gets all words that take part of the given lecture
        INPUT: leid: str: lecture ID
        OUTPUT: list of words (in french), phonetics and persian translations that are part of that lecture"""
        self.cursor.execute("SELECT * FROM Words WHERE leid = '{}' ".format(leid))
        # Take only the wanted results
        res = self.cursor.fetchall()
        res = [r[2:5] for r in res]
        return res

    def get_all_words(self):
        """Gets all words in the database
        OUTPUT: list of words (in french), phonetics and persian translations """
        # Find max lecture number
        self.cursor.execute("SELECT MAX(leid) FROM Lecture")
        last_id = self.cursor.fetchall()
        last_id = last_id[0][0]
        res = []
        for i in range(last_id+1):
            res.append(self.get_words_lecture(str(i)))
        # Format the result correctly
        res = [j for i in res for j in i]
        return res

    def search_word(self, word, language, leid='all'):
        """Search a word given the word and the language. The index of the table is set on the
            column 'leid' corresponding to the lecture ID.
            INPUT:  str: word = word to search for
                    str: language = language in which the word is given
                    int: leid = ID of the lecture -> speeds up search if given """
        # Select the word in the table
        if leid == 'all':
            self.cursor.execute("SELECT * FROM Words WHERE {} = '{}'".format(language, word))
            res = self.cursor.fetchall()
        else:
            self.cursor.execute("SELECT * FROM Words WHERE leid = '{}' AND {} = '{}'".format(leid, language, word))
            res = self.cursor.fetchall()
        return [i[2:5] for i in res]

    def get_id(self, word):
        """Gets all word ids corresponding to a word in french
        OUTPUT: list of ids """
        # Finds the lines matching that french word
        self.cursor.execute("SELECT * from Words where fran = '{}'".format(word))
        res = self.cursor.fetchall()
        print(res)
        return res[0][0]


    def import_images(self, file):
        """Adds the link of a word to an image
        INPUT: file containing the woid and an image name
        OUTPUT: update of the imag variable of the Words table to 'Y'
            addition of the image file to the Images table"""
        with open(file, 'r') as f:
            lines = f.readlines()
            for i in lines:
                i = i.rstrip()
                i = i.split(",")
                woid = i[0]
                # Check if image for the woid is already in table
                self.cursor.execute("SELECT * from Images where WOID = '{}'".format(woid))
                res = self.cursor.fetchall()
                if len(res) == 0:
                    # Add the image to the Imags table
                    self.cursor.execute("Insert into Images values ('{}', '{}')".format(woid, i[1]))
                    self.conn.commit()
                    # Set the boolean for image to 'Y'
                    self.cursor.execute("update Words set imag = 'Y' where woid = '{}'".format(woid))
                    self.conn.commit()

    def get_image(self, woid):
        # Get the image for this word
        self.cursor.execute("SELECT * from Images where WOID = '{}'".format(woid))
        res = self.cursor.fetchall()
        image_name = self.image_path + res[0][1] + ".jpeg"
        return image_name


handler = Handle_data(1)
#handler.add_lecture('0', '01/01/2000')
#handler.add_word(['lol', 'lol1', 'lol2'])
#handler.get_all_words()
#handler.get_random_word(1)
#handler.search_word('lol', 'fran', 1)
"""
handler.add_lecture('1', '04/10/2021')
handler.add_lecture('2', '11/10/2021')
handler.add_lecture('3', '18/10/2021')
handler.add_lecture('4', '25/10/2021')
handler.add_lecture('5', '08/11/2021')
handler.add_lecture('6', '15/11/2021')
handler.add_lecture('7', '22/11/2021')
handler.add_lecture('8', '29/11/2021')
handler.add_lecture('9', '06/12/2021')
handler.add_lecture('10', '13/12/2021')
handler.add_lecture('11', '20/12/2021')


handler.import_words('Data/Lecture1_1.txt')
handler.import_words('Data/Lecture1_2.txt')
handler.import_words('Data/Lecture1_3.txt')
handler.import_words('Data/Lecture1_4.txt')
handler.import_words('Data/Lecture1_5.txt')
handler.import_words('Data/Lecture1_6.txt')
handler.import_words('Data/Lecture1_7.txt')
handler.import_words('Data/Lecture1_8.txt')
handler.import_words('Data/Lecture1_9.txt')
handler.import_words('Data/Lecture1_10.txt')
handler.import_words('Data/Lecture1_11.txt')
"""


handler.import_images('/home/charlotte/Desktop/Projects/Persian/Data/images.txt')
#handler.get_image(2)

"""
handler.get_id('perroquet')
handler.get_id('seau')
handler.get_id('ticket')
handler.get_id('ligne')
handler.get_id('pluie')
handler.get_id('paysage')
handler.get_id('poupée')
handler.get_id('boîte')
handler.get_id('bougie')
handler.get_id('oie')
"""
