import commands.util
import sqlite3
import re
from constants import PY_SLACK, BOT_NAME
from commands.base_command import Base_Command

"""
Finds the doorcode if it exists for a particular query
"""
class Doorcode(Base_Command):
    def __init__(self, s):
        self.s = s
        self.match_patten = re.compile('[0-9]{3}')
        self.help_pattern = re.compile('-help')
        self.help_str = help_str = 'Produces the doorcode for a' \
        'specific class (three digit code), '\
        'just include \'doorcode\' or \'door code\' with a three '\
        'digit code in your request'

    def post_message(self, m):
        if self.help_pattern.search(self.s):
            commands.util.post(self.help_str, m)
        else:
            doorcode_str = self.post_doorcode(m)
            commands.util.post(doorcode_str, m)
        

    def post_doorcode(self, m):
        ret_str = ''
        nums = re.findall(self.match_patten, m.get('text'))
        if len(nums) < 1:
            return 'Sorry, your doorcode seems to be invalid'
        nums = nums[0]
        conn = sqlite3.connect('bot/commands/JARVIS')
        nums = (nums, )
        res = conn.execute('SELECT code FROM doorcode WHERE id=?',nums)
        ret = res.fetchone()
        if ret:
            ret_str = '`{0}`'.format(ret[0])
        else:
            ret_str = 'Sorry, I couldn\'t find a doorcode for {0}!'.format(nums[0])  
        return ret_str
   
        

       