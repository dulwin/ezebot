from constants import PY_SLACK, BOT_NAME

class Base_Command():
    def __init__(self, s):
        self.command = s
    def post_message(self, m):
        PY_SLACK.chat_post_message(
            m.get('channel'),
            'Hi there, <@{0}>!'.format(m.get('user')),
            username=BOT_NAME
        )
