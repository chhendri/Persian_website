from SQL_Handler import Handle_data
import random

class Controller:
    def __init__(self, seed):
        # Interactions with the database
        self.handler = Handle_data(seed)
        # Memory of the words and if they have been successfully translated by the
        self.memory = {}

    def get_random_word(self, lst):
        """Gets a random word given a list of words
        INPUT: lst: list of words"""
        return random.choice(lst)

    def choose_language(self, word, language):
        """Returns the word in a certain language"""
        if language == 'pers':
            w = word[2]
        elif language == 'fran':
            w = word[0]
        elif language == 'tran':
            w = word[1]
        return w

    def response_memory(self, given_r, corr_r):
        """Checks the given response and the correct response.
        INPUT:  str: given_r = translation given by the user
                str: corr_r = correct translation
        OUTPUT: dict: self.memory takes a new entry that gives the number of failures the user has"""
        if given_r == corr_r:
            self.memory[corr_r] = 0
            correct = True
        else:
            if corr_r in self.memory:
                self.memory[corr_r] += 1
            else:
                self.memory[corr_r] = 1
            correct = False
        return correct

    def ask_word(self, word_lst, l_ask, leid='all'):
        """Prompts a word and asks for the translation
        INPUT:  int: leid = lecture ID
                str: l_ask = language in which the word will be asked"""
        # Choose the language in which the result will be displayed
        if l_ask == 'pers':
            l_res = 'fran'
        elif l_ask == 'fran':
            l_res = 'pers'
        # Chooses a random word
        word = self.get_random_word(word_lst)
        # Choose a language to ask
        wq = self.choose_language(word, l_ask)
        print('What does ', wq, ' mean in ', l_res, ' ? ')
        return word, wq, l_ask, l_res, leid

    def treat_response(self, rr, word, wq, l_ask, l_res, leid):
        """Treats the response of the user"""
        # Correct response
        wr = self.handler.search_word(wq, l_ask, leid)
        wr2 = self.choose_language(wr[0], l_res)
        correct = self.response_memory(rr, wr2)
        if correct:
            print('Is correct!')
        else:
            print('Not correct!')
        print(wr[0])
        return word, correct

    def interrogation(self, l_ask='fran', leid='all'):
        """Prompts multiple words and prompts the incorrect words
        INPUT:
        str: input of the user: if the user inputs 'stop', stop the process.
                        if the user inputs 'skip', skip that word"""
        # Probability to change the language in which the question is asked
        proba_l = 0.5
        # Choose the list of words
        if leid == 'all':
            word_lst = self.handler.get_all_words()
        else:
            word_lst = self.handler.get_words_lecture(leid)
        # Pool of words to choose from
        word_pool = word_lst
        inp = 'run'
        while inp == 'run':
            # Ask a question
            word, wq, l_ask, l_res, leid = self.ask_word(word_lst, l_ask, leid)
            # Ask for input from the user
            rr = str(input('The response is : '))
            if rr != 'skip':
                # Check the result of the question
                asked = self.treat_response(rr, word, wq, l_ask, l_res, leid)
                if asked[1]:
                    word_pool.remove(asked[0])
            if rr == 'stop' or word_pool == []:
                inp = 'stop'
                print('Thanks for using our software!')




c = Controller(3)
c.interrogation(leid=5)
