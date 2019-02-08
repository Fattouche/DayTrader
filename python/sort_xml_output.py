import xml.etree.ElementTree as ET

# def sortchildrenby(parent, attr):
#     parent[:] = sorted(parent, key=lambda child: child.get(attr))

tree = ET.parse('testLOG')
root = tree.getroot()

# child[0] is the timestamp
root[:] = sorted(root, key=lambda child: child[0].text)

tree.write('output.xml')