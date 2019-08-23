from __future__ import unicode_literals

from django.conf import settings
from django.http import HttpResponse
from django.shortcuts import render
from django.template.loader import render_to_string

from .models import Tank


def index(request):
    return render(request, "tanks/index.html")

def forf_manual(request):
    return render(request, "tanks/forf.html")

def procedure_reference(request):
    return render(request, "tanks/procs.html")

def debugger(request):
    return render(request, "tanks/debugger.html")

def upload(request):
    token = request.POST.get("token")
    program = request.POST.get("program")
    tank_name = request.POST.get("name")
    author = request.POST.get("author")
    color = request.POST.get("color")
    sensors = []
    for i in range(10):
        sensor = {}
        sensor["range"] = request.POST.get("s%sr" % (i,))
        sensor["angle"] = request.POST.get("s%sa" % (i,))
        sensor["width"] = request.POST.get("s%sw" % (i,))
        sensor["is_turret"] = True if request.POST.get("s%st" % (i,))=="on" else False
        sensors.append(sensor)

    tank = {"token": token, "program": program, "tank_name": tank_name, "author": author, "color": color, "sensors": sensors}
    print(tank)
        
    return HttpResponse("", status=200)

def state(request, token):
    return HttpResponse("", status=200)
