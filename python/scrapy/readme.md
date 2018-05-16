    pip install scrapy

    scrapy startproject py123demo
    cd py123demo
    scrapy genspider demo python123.io
    
    edit spiders/demo.py
    
    scrapy crawl demo
    