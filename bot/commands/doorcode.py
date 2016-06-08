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
        self.match_patten = '[0-9]{3}'

    def post_message(self, m):
        ret_str = ''
        nums = re.findall(self.match_patten, m.get('text'))
        nums = nums[0]
        conn = sqlite3.connect('bot/commands/JARVIS')
        nums = (nums, )
        res = conn.execute('SELECT code FROM doorcode WHERE id=?',nums)
        ret = res.fetchone()
        if ret:
            ret_str = '`{0}`'.format(ret[0])
        else:
            ret_str = 'Sorry, I couldn\'t find a doorcode for {0}!'.format(nums[0])  
        commands.util.post(ret_str, m)

       