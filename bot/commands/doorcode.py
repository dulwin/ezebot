from constants import PY_SLACK, BOT_NAME
import sqlite3
import re

class Doorcode():
    def __init__(self, s):
        self.s = s
        self.match_patten = '[0-9]{3}'

    def post_message(self, m):
        nums = re.findall(self.match_patten, m.get('text'))[0]
        conn = sqlite3.connect('bot/commands/JARVIS')
        nums = (nums, )
        res = conn.execute('SELECT code FROM doorcode WHERE id=?',nums)
        ret = res.fetchone()
        if ret:
            ret_str = '`{0}`'.format(ret[0])
        else:
            ret_str = 'Sorry, I couldn\'t find a doorcode for {0}!'.format(nums[0])    

        PY_SLACK.chat_post_message(
            m.get('channel'),
            ret_str,
            username=BOT_NAME
        )