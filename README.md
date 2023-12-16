**به زودی اپدیت میشود**

![R (2)](https://github.com/Azumi67/PrivateIP-Tunnel/assets/119934376/a064577c-9302-4f43-b3bf-3d4f84245a6f)
نام پروژه : تانل UDP2RAW + FEC[UDPSPEED] برای وایرگارد
---------------------------------------------------------------

-----------
**توضیح کوتاه در مورد این پروژه :**

- حتما در سرور تست، نخست تانل را ازمایش کنید و سپس اقدام به استفاده از آن بکنید.
- تمامی جوانب را در نظر بگیرید و از تانل استفاده نمایید.
- باید توجه داشته باشید برای تمرین و اموزش خودم، اقدام به ساخت اسکریپت میکنم و در کنارش آموزش هم مینویسم که شما اگر خواستید استفاده کنید.
- این اسکریپت با زبان GO نوشته شده است و از انجا که در مسیر یادگیری هستم، اگر مشکلی یا کاستی در آن دیدید، ببخشید.
- در این اسکریپت از منوی جدیدی استفاده کردم. میتوانید با کیبورد، گزینه مورد نظر را انتخاب کنید و سپس ENTER بزنید.
- در این تانل شما میتوانید از ICMP[IPV4] و IP6IP6 و FEC و بدون FEC استفاده کنید.
- هیچ پورت دیفالتی در این تانل گذاشته نشده است.
- پنل وایرگارد در خارج باید نصب شده باشد یا اگر بدون پنل هستید ، باید وایرگارد در خارج نصب شده باشد.
- لطفا برای کانفیگ دوباره، نخست از منوی uninstall اقدام به حذف تانل کنید تا مشکلی پیش نیاید.
- در آخر هر کانفیگ، ایپی 4 سرور ایران شما با پورت نهایی نمایش داده میشود.
- من در وقت آزاد این را درست کردم و ممکن است اشتباهاتی هم داخلش باشد. پیشاپیش ببخشید.
--------------

![Exclamation-Mark-PNG-Clipart](https://github.com/Azumi67/Game_tunnel/assets/119934376/3951d7d9-0e17-4723-b07f-786500ccbc7f)**چند نکته**

- برای تانل ICMP ، حتما اگر اشتباهی در کانفیگ انجام دادید باید حتما هم در سرور ایران و خارج حذفش کنید و هر دو سرور ریبوت شود در غیر این صورت خطای SERVER IS FULL را میگیرید.
- قبل از کانفیگ دوباره، همیشه با دستور ip a مشاهده کنید که tun0 یا tun1 که مربوط به icmp است ، موجود نباشد. حتما پس از Uninstall ICMP سرور خود را ریست نمایید.
- مورد دیگر اینکه، در سرور های ایران اگر DNS مشکل داشته باشد، ممکن است دانلود انجام نشود. حتما از طریق nano /etc/resolv.conf اقدام به تغییر موقتی dns خود بکنید .
- ممکن است در سرور ایران شما، سرعت دانلود پایین باشد و برای همین، ممکنه که دانلود پیش نیاز ها کمی طول بکشد.
- پورت ها در آموزش برای مثال استفاده شده اند، شما میتوانید از پورت های دلخواه خودتان استفاده نمایید.


------------------------
![307981](https://github.com/Azumi67/V2ray_loadbalance_multipleServers/assets/119934376/39b2794b-fd04-4ae5-baea-d4b66138766e)
 **فهرست :**
 ----------------------
 
**دسترسی سریع به اسکریپت** >> **[کلیک]()**
- **امکانات** >> **[کلیک](https://github.com/Azumi67/UDP2RAW_FEC#%D8%A7%D9%85%DA%A9%D8%A7%D9%86%D8%A7%D8%AA)**
- **تک سرور** : 
- **تانل UDP2RAW FEC IPV4** >> **[کلیک](https://github.com/Azumi67/UDP2RAW_FEC#%D8%AA%D8%A7%D9%86%D9%84-udp2raw-fec-ipv4-%D8%AA%DA%A9-%DA%A9%D8%A7%D9%86%D9%81%DB%8C%DA%AF)**
- **تانل UDP2RAW FEC IPV6** >> **[کلیک](https://github.com/Azumi67/UDP2RAW_FEC#%D8%AA%D8%A7%D9%86%D9%84-udp2raw-fec-ipv4-%D8%AA%DA%A9-%DA%A9%D8%A7%D9%86%D9%81%DB%8C%DA%AF-1)**
- **تانل UDP2RAW FEC ICMP** >> **[کلیک](https://github.com/Azumi67/UDP2RAW_FEC#%D8%AA%D8%A7%D9%86%D9%84-udp2raw-fec-icmp-%D8%AA%DA%A9-%DA%A9%D8%A7%D9%86%D9%81%DB%8C%DA%AF)**
- **تانل UDP2RAW FEC + PrivateIP** >>  **[کلیک](https://github.com/Azumi67/UDP2RAW_FEC#%D8%AA%D8%A7%D9%86%D9%84-udp2raw-fec--privateip)**
- **اسکریپت های کارآمد** >> **[کلیک](https://github.com/Azumi67/UDP2RAW_FEC#%D8%A7%D8%B3%DA%A9%D8%B1%DB%8C%D9%BE%D8%AA-%D9%87%D8%A7%DB%8C-%DA%A9%D8%A7%D8%B1%D8%A2%D9%85%D8%AF-)**

--------------
 
 ![Exclamation-Mark-PNG-Clipart](https://github.com/Azumi67/KCP_tunnel/assets/119934376/96bc9a31-57d1-4c14-a988-f9284f547955)**نکته**

- من تنها وقت تست تک کانفیگ وایرگارد را داشتم. شما میتوانید چندین کانفیگ را تست کنید و به من فیدبک بدید.

------------------------
![check](https://github.com/Azumi67/PrivateIP-Tunnel/assets/119934376/13de8d36-dcfe-498b-9d99-440049c0cf14)
**امکانات**
-

- پشتیبانی از UDP
- پایین اوردن PACKET LOST در کانکشن نهایی
- قابلیت تانل بر روی تک پورت و 5 پورت
- امکان استفاده از ایپی 4 و 6
- استفاده از FEC در تانل
- امکان استفاده از IP6IP6 و تانل UDP2RAW به همراه FEC
- امکان استفاده ار ICMP با پرایوت ایپی 4 و تانل UDP2RAW به همراه FEC
- ایجاد سرویس برای تمامی گزینه ها
- امکان حذف تمامی تانل ها و سرویس ها

-------------------

What is FEC ?
FEC stands for Forward Error Correction. It is a technique used in data communication to enhance the reliability of data transmission over unreliable or noisy channels. The purpose of FEC is to detect and correct errors that may occur during transmission without the need for retransmission.
Overall, FEC helps improve the reliability and quality of data transmission by adding error correction codes to the transmitted data, allowing for the detection and correction of errors at the receiving end.

 ------------------------------------------------------

![147-1472495_no-requirements-icon-vector-graphics-clipart](https://github.com/Azumi67/V2ray_loadbalance_multipleServers/assets/119934376/98d8c2bd-c9d2-4ecf-8db9-246b90e1ef0f)
 **پیش نیازها**

 - لطفا سرور اپدیت شده باشه.
 - میتوانید از اسکریپت اقای [Hwashemi](https://github.com/hawshemi/Linux-Optimizer) و یا [OPIRAN](https://github.com/opiran-club/VPS-Optimizer) هم برای بهینه سازی سرور در صورت تمایل استفاده نمایید. (پیش نیاز نیست)


----------------------------

  
  ![6348248](https://github.com/Azumi67/PrivateIP-Tunnel/assets/119934376/398f8b07-65be-472e-9821-631f7b70f783)
**آموزش**
-
![OIP2 (1)](https://github.com/Azumi67/V2ray_loadbalance_multipleServers/assets/119934376/3ec2f05f-3308-4441-8cce-62ab4776f4e2)
**تانل UDP2RAW FEC IPV4**
----------------------------------
![green-dot-clipart-3](https://github.com/Azumi67/6TO4-PrivateIP/assets/119934376/902a2efa-f48f-4048-bc2a-5be12143bef3) **سرور خارج**

**مسیر : UDP2RAW FEC IPV4 > Kharej**



 <p align="right">
  <img src="https://github.com/Azumi67/UDP2RAW_FEC/assets/119934376/a2da63d2-96a9-45b1-95b4-da0e73159f4b" alt="Image" />
</p>


- نخست سرور خارج را کانفیگ میکنیم
- خب پیش نیاز ها در صورت نیاز دانلود میشوند.
- تعداد کانفیگ را عدد 1 وارد میکنم چون تنها یک کانفیگ دارم
- پورت تانل را 443 قرار میدم
- پورت FEC را 3333 قرار میدم. دقت نمایید این پورت نهایی شما خواهد بود.
- پسورد را azumi قرار میدم
- پورت وایرگارد من در سرور خارج 20820 میباشد.
- من raw-mode را برای مثال icmp انتخاب میکنم.
- حالا باید سرور ایران را کانفیگ کرد.
----------------------

![green-dot-clipart-3](https://github.com/Azumi67/6TO4-PrivateIP/assets/119934376/902a2efa-f48f-4048-bc2a-5be12143bef3) **سرور ایران** 

**مسیر : UDP2RAW FEC IPV4 > IRAN**




 <p align="right">
  <img src="https://github.com/Azumi67/UDP2RAW_FEC/assets/119934376/5611343d-a79e-4bc2-b4d4-e1791f0bddbb" alt="Image" />
</p>

- سپس سرور ایران را کانفیگ میکنیم
- پیش نیاز ها به صورت اتوماتیک در صورت AVAILABLE نبودن؛ دانلود خواهند شد.
- ایپی 4 سرور خارج را وارد نمایید
- پورت تانل را در سرور خارج 443 وارد کردیم.
- پورت FEC هم در سرور خارج 3333 وارد کردیم. این پورت نهایی ما خواهد بود.
- پسورد هم که در سرور خارج azumi وارد کردیم.
- در سرور خارج raw mode را icmp انتخاب کرده بودیم.
- در آخر ایپی سرور ایران شما با پورت نهایی برای وایرگارد نمایش داده میشود.

-------------------
![Exclamation-Mark-PNG-Clipart](https://github.com/Azumi67/UDP2RAW_FEC/assets/119934376/270e7fa5-6b7d-472c-b3ce-2f982a2f0cee)**نکته**

- برای تانل udp2raw بدون fec اموزشی قرار ندادم چون همه دیگه بهش اشنا هستید.
- اما به صورت کلی پورت تانل و پورت وایرگارد و پسورد و raw mode را انتخاب میکنید.
- و در سرور ایران هم مانند سرور خارج، تمام موارد بالا به اضافه ایپی خارج را وارد میکنید.

--------------------------------------
![OIP2 (1)](https://github.com/Azumi67/V2ray_loadbalance_multipleServers/assets/119934376/3ec2f05f-3308-4441-8cce-62ab4776f4e2)
**تانل UDP2RAW FEC IPV6**
-
![green-dot-clipart-3](https://github.com/Azumi67/6TO4-PrivateIP/assets/119934376/902a2efa-f48f-4048-bc2a-5be12143bef3) **سرور خارج**

**مسیر : UDP2RAW FEC IPV6 > Kharej**


 <p align="right">
  <img src="https://github.com/Azumi67/UDP2RAW_FEC/assets/119934376/57f51289-1e48-4659-b423-f8b12e64cf0e" alt="Image" />
</p>

- نخست سرور خارج را کانفیگ میکنیم
- خب پیش نیاز ها در صورت نیاز دانلود میشوند.
- تعداد کانفیگ را عدد 1 وارد میکنم چون تنها یک کانفیگ دارم
- پورت تانل را 443 قرار میدم
- پورت FEC را 3333 قرار میدم. دقت نمایید این پورت نهایی شما خواهد بود.
- پسورد را azumi قرار میدم
- پورت وایرگارد من در سرور خارج 20820 میباشد.
- من raw-mode را برای مثال icmp انتخاب میکنم.
- حالا باید سرور ایران را کانفیگ کرد.
----------------------

![green-dot-clipart-3](https://github.com/Azumi67/6TO4-PrivateIP/assets/119934376/902a2efa-f48f-4048-bc2a-5be12143bef3) **سرور ایران** 

**مسیر : UDP2RAW FEC IPV6 > IRAN**


 <p align="right">
  <img src="https://github.com/Azumi67/UDP2RAW_FEC/assets/119934376/5273fb28-cf03-49ef-8a94-208466282e93" alt="Image" />
</p>

- سپس سرور ایران را کانفیگ میکنیم
- پیش نیاز ها به صورت اتوماتیک در صورت AVAILABLE نبودن؛ دانلود خواهند شد.
- ایپی 6 سرور خارج را وارد نمایید
- پورت تانل را در سرور خارج 443 وارد کردیم.
- پورت FEC هم در سرور خارج 3333 وارد کردیم. این پورت نهایی ما خواهد بود.
- پسورد هم که در سرور خارج azumi وارد کردیم.
- در سرور خارج raw mode را icmp انتخاب کرده بودیم.
- در آخر ایپی سرور ایران شما با پورت نهایی برای وایرگارد نمایش داده میشود.
--------------------------------------
![OIP2 (1)](https://github.com/Azumi67/V2ray_loadbalance_multipleServers/assets/119934376/3ec2f05f-3308-4441-8cce-62ab4776f4e2)
**تانل UDP2RAW FEC ICMP**
----------------------------------
![green-dot-clipart-3](https://github.com/Azumi67/6TO4-PrivateIP/assets/119934376/902a2efa-f48f-4048-bc2a-5be12143bef3) **سرور خارج**

**مسیر : UDP2RAW FEC ICMP > KHAREJ**



 <p align="right">
  <img src="https://github.com/Azumi67/UDP2RAW_FEC/assets/119934376/1edd6567-17ef-48e7-8b4e-4c0a689ed404" alt="Image" />
</p>

-**باید برای کانفیگ دوباره حتما کانفیگ قدیمی را uninstall کنید.**
- نخست سرور خارج را کانفیگ میکنیم
- اگر میخواهید توسط پرایوت ایپی  4 و تانل icmp ، تانل kcp را برقرار کنید، این روش برای شما مناسب است.
- حتما دقت نمایید که قبلا این تانل را نساخته باشید چون دیوایس جدید برای شما میسازد. پس حتما با دستور ip a از موجود نبودن آن اطمینان حاصل فرمایید.
- در صورت موجود بودن آن حتما اقدام به حذف آن نمایید و سپس سرور خود را ریبوت کنید و سپس اقدام به کانفیگ دوباره نمایید.
- خب پیش نیاز ها در صورت نیاز دانلود میشوند.
- تعداد کانفیگ را عدد 1 وارد میکنم چون تنها یک کانفیگ دارم
- پورت تانل را 443 قرار میدم
- پورت FEC را 3333 قرار میدم. دقت نمایید این پورت نهایی شما خواهد بود.
- پسورد را azumi قرار میدم
- پورت وایرگارد من در سرور خارج 20820 میباشد.
- من raw-mode را برای مثال udp انتخاب میکنم.
- حالا باید سرور ایران را کانفیگ کرد.
----------------------

![green-dot-clipart-3](https://github.com/Azumi67/6TO4-PrivateIP/assets/119934376/902a2efa-f48f-4048-bc2a-5be12143bef3) **سرور ایران** 

**مسیر : UDP2RAW FEC ICMP > IRAN**

 <p align="right">
  <img src="https://github.com/Azumi67/UDP2RAW_FEC/assets/119934376/9054989e-b1e7-4d57-8cb1-74156d496194" alt="Image" />
</p>

- سپس سرور ایران را کانفیگ میکنیم
- پس از نصب icmptunnel، ایپی 4 سرور خارج خودتان را وارد نمایید.
- سپس اگر مراحل را درست رفته باشید باید تانل icmp شما برقرار شده باشد.
- پیش نیاز ها به صورت اتوماتیک در صورت AVAILABLE نبودن؛ دانلود خواهند شد.
- پورت تانل را در سرور خارج 443 وارد کردیم.
- پورت FEC هم در سرور خارج 3333 وارد کردیم. این پورت نهایی ما خواهد بود.
- پسورد هم که در سرور خارج azumi وارد کردیم.
- در سرور خارج raw mode را udp انتخاب کرده بودیم.
- در آخر ایپی سرور ایران شما با پورت نهایی برای وایرگارد نمایش داده میشود.

--------------------------------------
![OIP2 (1)](https://github.com/Azumi67/V2ray_loadbalance_multipleServers/assets/119934376/3ec2f05f-3308-4441-8cce-62ab4776f4e2)
**تانل UDP2RAW FEC + PrivateIP**
----------------------------------
![green-dot-clipart-3](https://github.com/Azumi67/6TO4-PrivateIP/assets/119934376/902a2efa-f48f-4048-bc2a-5be12143bef3) **سرور خارج**

**مسیر : UDP2RAW FEC IP6IP6 > KHAREJ**



 <p align="right">
  <img src="https://github.com/Azumi67/UDP2RAW_FEC/assets/119934376/79fc3b6b-3391-4c1f-801d-881e6d6e3a9" alt="Image" />
</p>

**قبل از کانفیگ ، اطمینان پیدا کنید که تانل 6to4 یا تانل های دیگری ندارید که خطای بافر سایز نگیرید**
- نخست سرور خارج را کانفیگ میکنیم
- میخواهیم از طریق IP6IP6 و UDP2RAW همراه با FEC، تانل را برقرار کنیم.
- حتما دقت نمایید که قبلا پرایوت ایپی نداشته باشید که خطای بافر سایز نگیرید.
- پس از حذف کردن پرایوت ایپی ، حتما یک بار ریبوت بفرمایید.
- ایپی 4 سرور خارج و ایران را میدهم.
- خب حالا نوبت کانفیگ تانل میباشد.
- خب پیش نیاز ها در صورت نیاز دانلود میشوند.
- تعداد کانفیگ را عدد 1 وارد میکنم چون تنها یک کانفیگ دارم
- پورت تانل را 443 قرار میدم
- پورت FEC را 3333 قرار میدم. دقت نمایید این پورت نهایی شما خواهد بود.
- پسورد را azumi قرار میدم
- پورت وایرگارد من در سرور خارج 20820 میباشد.
- من raw-mode را برای مثال icmp انتخاب میکنم.
- حالا باید سرور ایران را کانفیگ کرد.

----------------------

![green-dot-clipart-3](https://github.com/Azumi67/6TO4-PrivateIP/assets/119934376/902a2efa-f48f-4048-bc2a-5be12143bef3) **سرور ایران** 

**مسیر : UDP2RAW FEC IP6IP6 > IRAN**


 <p align="right">
  <img src="https://github.com/Azumi67/UDP2RAW_FEC/assets/119934376/689d5498-acb7-4d6c-81dc-48ecb91b3227" alt="Image" />
</p>

- سرور ایران را کانفیگ میکنیم
- ایپی 4 سرور خارج و ایران را میدهم.
- سپس برای شما پینگ میگیرد.
- سپس تانل UDP2RAW + FEC را کانفیگ میکنیم
- پیش نیاز ها به صورت اتوماتیک در صورت AVAILABLE نبودن؛ دانلود خواهند شد.
- پورت تانل را در سرور خارج 443 وارد کردیم.
- پورت FEC هم در سرور خارج 3333 وارد کردیم. این پورت نهایی ما خواهد بود.
- پسورد هم که در سرور خارج azumi وارد کردیم.
- در سرور خارج raw mode را icmp انتخاب کرده بودیم.
- در آخر ایپی سرور ایران شما با پورت نهایی برای وایرگارد نمایش داده میشود.

--------------------------------------
**اسکرین شات**


<details>
  <summary align="right">Click to reveal image</summary>
  
  <p align="right">
    <img src="https://github.com/Azumi67/UDP2RAW_FEC/assets/119934376/3e1ed4b1-53f3-4aec-b5bd-ddeda59577ba" alt="menu screen" />
  </p>
</details>


------------------------------------------
![scri](https://github.com/Azumi67/FRP-V2ray-Loadbalance/assets/119934376/cbfb72ac-eff1-46df-b5e5-a3930a4a6651)
**اسکریپت های کارآمد :**
-
- این اسکریپت ها optional میباشد.


 
 Opiran Scripts
 
```
 bash <(curl -s https://raw.githubusercontent.com/opiran-club/pf-tun/main/pf-tun.sh --ipv4)
```

```
apt install curl -y && bash <(curl -s https://raw.githubusercontent.com/opiran-club/VPS-Optimizer/main/optimizer.sh --ipv4)
```

Hawshemi script

```
wget "https://raw.githubusercontent.com/hawshemi/Linux-Optimizer/main/linux-optimizer.sh" -O linux-optimizer.sh && chmod +x linux-optimizer.sh && bash linux-optimizer.sh
```

-----------------------------------------------------
![R (a2)](https://github.com/Azumi67/PrivateIP-Tunnel/assets/119934376/716fd45e-635c-4796-b8cf-856024e5b2b2)
**اسکریپت من**
----------------

```

```

- اپدیت میشود

```

```

 <div dir="rtl">&bull;  دستور زیر برای کسانی هست که پیش نیاز ها را در سرور، نصب شده دارند</div>
 
```

```

---------------------------------------------
![R (7)](https://github.com/Azumi67/PrivateIP-Tunnel/assets/119934376/42c09cbb-2690-4343-963a-5deca12218c1)
**تلگرام** 
![R (6)](https://github.com/Azumi67/FRP-V2ray-Loadbalance/assets/119934376/f81bf6e1-cfed-4e24-b944-236f5c0b15d3) [اپیران- OPIRAN](https://github.com/opiran-club)

---------------------------------
![R23 (1)](https://github.com/Azumi67/FRP-V2ray-Loadbalance/assets/119934376/18d12405-d354-48ac-9084-fff98d61d91c)
**سورس ها**




![R (9)](https://github.com/Azumi67/FRP-V2ray-Loadbalance/assets/119934376/33388f7b-f1ab-4847-9e9b-e8b39d75deaa) [سورس  OPIRAN](https://github.com/opiran-club)

![R (9)](https://github.com/Azumi67/6TO4-GRE-IPIP-SIT/assets/119934376/4758a7da-ab54-4a0a-a5a6-5f895092f527)[سورس  Hwashemi](https://github.com/hawshemi/Linux-Optimizer)



-----------------------------------------------------

![youtube-131994968075841675](https://github.com/Azumi67/FRP-V2ray-Loadbalance/assets/119934376/24202a92-aff2-4079-a6c2-9db14cd0ecd1)
**ویدیوی آموزش**

-----------------------------------------
