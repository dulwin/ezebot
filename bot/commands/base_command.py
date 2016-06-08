from constants import PY_SLACK, BOT_NAME
import commands.util

"""
The idea here is that all commands inherit from base command.
This way the consumer program can use a polymorphic class and not really care about
the actual result of the regex
"""
class Base_Command():
    def __init__(self, s):
        self.command = s
    def post_message(self, m):
        msg = 'Hi there, <@{0}>!'.format(m.get('user'))
        commands.util.post(msg, m)
