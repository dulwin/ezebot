import re
import os
from commands.doorcode import Doorcode
from commands.base_command import Base_Command

doorcode_pattern = re.compile('doorcode.*[0-9]')

def find_match(s):
    if doorcode_pattern.search(s):
        return Doorcode(s)
    else:
        return Base_Command(s)