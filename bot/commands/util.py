import re
import os
from commands.doorcode import Doorcode
from commands.base_command import Base_Command
from constants import PY_SLACK, BOT_NAME

doorcode_pattern = re.compile('doorcode.*[0-9]{3}')

"""
Until better design is created, this function is designed 
to take a string and run it through a variety of regex matches before deciding what
class it belongs to
"""
def find_match(s):
    if doorcode_pattern.search(s):
        return Doorcode(s)
    else:
        return Base_Command(s)


"""
Helper to post messages to slack, posts to channel of original message
"""
def post(s, m):
     PY_SLACK.chat_post_message(
            m.get('channel'),
            s,
            username=BOT_NAME
        )