from flask import Flask, request, jsonify
from hugchat import hugchat
from hugchat.login import Login

EMAIL = 'grechin3003@gmail.com'
PASSWORD = 'HellBoii2033'
cookie_path_dir = "./cookies/" 

sign = Login(EMAIL, PASSWORD)
cookies = sign.login(cookie_dir_path=cookie_path_dir, save_cookies=True)
chatbot = hugchat.ChatBot(cookies=cookies.get_dict())

app = Flask(__name__)

@app.route('/ask', methods=['POST'])
def ask():
    user_input = request.json.get('text')
    response = generate_answ(user_input)
    return jsonify({'response': response})

def generate_answ(text):
    prompt = text

    prompt_llm = f"""
    System:
    Когда к тебе обратяться для объяяснение что это за зона, расскажи что в ней опасного, почему тут нельзя загрязнять или строить что-то.
    При строительстве или работе в природных зонах учитывайте:
    Экологическое воздействие: Изменения могут повредить экосистему, снизить биоразнообразие и ухудшить среду обитания.
    Загрязнение: Строительные работы могут загрязнить воду, воздух и почву, разрушить растительность.
    Влияние на животных: Нарушение мест обитания и взаимодействие с животными могут вызвать стресс и гибель.
    Законодательство: Соблюдайте законы, защищающие окружающую среду, чтобы избежать штрафов.
    Этика: Уважение к природе и дикой жизни — наша ответственность.
    Соблюдение этих принципов помогает сохранить природное наследие и баланс экосистем.

    User:
    {prompt}

    Answer:
    """

    mes_res = chatbot.chat(prompt_llm)
    return mes_res.text

if __name__ == '__main__':
    app.run(port=5000)
    