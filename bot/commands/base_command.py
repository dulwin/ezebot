from constants import PY_SLACK, BOT_NAME

"""
The idea here is that all commands inherit from base command.
This way the consumer program can use a polymorphic class and not really care about
the actual result of the regex
"""
class Base_Command():
    def __init__(self, s):
        self.command = s
    def post_message(self, m):
        PY_SLACK.chat_post_message(
            m.get('channel'),
            'Hi there, <@{0}>!'.format(m.get('user')),
            username=BOT_NAME
        )
