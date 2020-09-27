import selenium as se
from selenium import webdriver
import time
#导入库
base_url="https://newsso.shu.edu.cn/login"
driver = webdriver.Chrome()
driver.get(base_url)
#使用之前需要安装selenium库并安装chorme驱动
username=''
password=''#你需要填写的账号和密码

name_input = driver.find_element_by_name('username')

pass_input = driver.find_element_by_name('password')
name_input.clear()
name_input.send_keys(username)  # 填写账号

pass_input.clear()
pass_input.send_keys(password) 
driver.find_element_by_id('login-submit').click()
#登录每日一报
base_url="https://selfreport.shu.edu.cn/XueSFX/FanXRB.aspx"
driver.get(base_url)
#跳转至手机端填报窗口
day_report=driver.find_element_by_id('p1_Button1')
day_report.click()
driver.find_element_by_id('p1_ChengNuo-inputEl-icon').click()
day_temp=driver.find_element_by_id('p1_TiWen-inputEl')
day_temp.clear()
day_temp.send_keys('36')  # 温度
driver.find_element_by_id('fineui_7-inputEl-icon').click()#绿色

time.sleep(1)
driver.find_element_by_id('p1_ctl00_btnSubmit').click() #提交
time.sleep(2)
driver.find_element_by_id('fineui_14').click() #提交确认
time.sleep(3)
#晨报
driver.close()
base_url="https://selfreport.shu.edu.cn/XueSFX/FanXRB.aspx"
driver.get(base_url)

day_report=driver.find_element_by_id('p1_Button2')
day_report.click()
driver.find_element_by_id('p1_ChengNuo-inputEl-icon').click()
day_temp=driver.find_element_by_id('p1_TiWen-inputEl')
day_temp.clear()
day_temp.send_keys('36')  # 温度
driver.find_element_by_id('fineui_7-inputEl-icon').click()#绿色

time.sleep(1)
driver.find_element_by_id('p1_ctl00_btnSubmit').click() #提交
time.sleep(2)
driver.find_element_by_id('fineui_14').click() #提交确认
time.sleep(3)
#晚报
driver.close()