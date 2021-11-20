import socket, glob, json

port = 53
ip = '127.0.1'

sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
sock.bind((ip, port))

def load_zones():

    jsonzone = {}
    zonefiles = glob.glob('*.zone')

    for zonefile in zonefiles:
        with open(zonefile, 'r') as f:
            data = json.load(f)
            zonename = data["$origin"]
            jsonzone[zonename] = data
    return jsonzone

zonedata = load_zones()

def getflags(flags):

    byte1 = bytes([flags[:1]])
    byte2 = bytes([flags[1:2]])

    rflag = ''

    QR = ''
    QRCODE = ''
    for bit in range(1, 5):
        QRCODE += str(ord(byte1) & (1 << bit))

    AA = '1'
    TC = '0'
    RD = '0'
    RA = '0'
    Z = '000'
    RCODE = '0000'

    return int(QRCODE + AA + TC + RD + RA + Z + RCODE, 2).to_bytes(1, byteorder='big')+int(byte2, 2).to_bytes(1, byteorder='big')

def getquestiondomain(data):

    state = 0
    exceptedlength = 0
    domain = ''
    domainparts = []
    x = 0
    y = 0
    for byte in data:
        if state == 0:
            if byte == 0:
                state = 1
                exceptedlength = 0
            else:
                exceptedlength += 1
        elif state == 1:
            if byte == 0:
                domainparts.append(domain)
                domain = ''
                state = 0
            else:
                domain += chr(byte)
        else:
            print('error')
    return domainparts

def getquestiontype(data):

    type = ''
    for byte in data:
        type += chr(byte)
    return type

def getzone(domain):

    for zone in zonedata:
        if domain.endswith(zone):
            return zone
    return None

def getrecs(data):

    recs = []
    for rec in data:
        recs.append(rec)
    return recs

def buildquestions(domainname, rectype):
    
        domainparts = getquestiondomain(domainname)
        domainparts.append(rectype)
        domainparts.reverse()
        domainparts = [len(part).to_bytes(1, byteorder='big') + part.encode() for part in domainparts]
        domainparts = b''.join(domainparts)
        return domainparts

def recttype(domainname, rectype, rectll, recval):

    rbytes = b''

    if rectll == 'a':
        rbytes = recval.split('.')
    return rbytes

def buildresponse(data):

    #transaction id
    tid = data[:2]

    #get flags
    flags = data[2:4]

    #question count
    qcount = data[4:6]

    #answer count
    acount = data[6:8]

    #name server count
    nscount = data[8:10]

    #additional record count
    arcount = data[10:12]

    dnsheader = tid + flags + qcount + acount + nscount + arcount

    #create dns body
    dnsbody = b''

    #get answer for query
    domainname = data[12:]
    rectype = getquestiontype(data[12:])
    zone = getzone(domainname)
    if zone is not None:
        recs = getrecs(zonedata[zone]['$origin'])
        for rec in recs:
            if rec['type'] == rectype:
                dnsbody += buildquestions(domainname, rectype)
                dnsbody += recttype(domainname, rectype, rec['type'], rec['value'])
                break
    else:
        dnsbody = b'\x00\x00\x00\x00'

    return dnsheader + dnsbody

while True:
    data, addr = sock.recvfrom(1024)
    print(addr)
    print(data)
    resp = buildresponse(data)
    sock.sendto(resp, addr)
