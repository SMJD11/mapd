import urllib.request
import json
from string import Template

url_base = "https://maps.planogis.org/arcgiswad/rest/services/Engineering/SafeRoutesToSchool_Elementary/MapServer/5/query?where=1%3D1&outFields=*&outSR=4326&f=geojson"

print("Downloading Stop Signs from ArcGIS...")

features = []
offset = 0
limit = 2000

while True:
    print(f"Fetching from offset {offset}...")
    url = f"{url_base}&resultOffset={offset}&resultRecordCount={limit}"
    response = urllib.request.urlopen(url)
    data = json.loads(response.read())
    
    batch = data.get("features", [])
    if not batch:
        break
        
    features.extend(batch)
    if len(batch) < limit:
        break
    offset += limit

print(f"Downloaded {len(features)} total stop signs.")

osm_xml = """<?xml version='1.0' encoding='UTF-8'?>
<osm version="0.6" generator="ArcGIS to OSM Convertor">
"""

osm_node_template = Template("""
  <node id="-$id" lat="$lat" lon="$lon" visible="true">
    <tag k="highway" v="stop" />
    <tag k="direction" v="$direction" />
    <tag k="note" v="$text" />
  </node>""")

# OSM Node IDs must be unique and negative for JOSM to treat them as "new nodes to insert".
for idx, feature in enumerate(features):
    props = feature["properties"]
    geom = feature["geometry"]["coordinates"]
    
    # The ArcGIS mapping gives us the exact physical location of the Stop Sign pole.
    # We will use highway=stop here, but when importing to OSM, these nodes MUST be dragged onto
    # the existing roads (ways) manually in JOSM for routing engines to understand them.
    travel_dir = props.get("TRAVELDIRECTION", "")
    
    node_str = f'  <node id="-{idx + 1}" lat="{geom[1]}" lon="{geom[0]}" visible="true" action="modify">\n'
    node_str += f'    <tag k="highway" v="stop" />\n'
    
    if travel_dir and travel_dir.strip() != "":
        # Map the general cardinal direction to standard OSM traffic_sign orientation tagging
        node_str += f'    <tag k="direction" v="{travel_dir}" />\n'
        
    if "SIGNTEXT" in props and props["SIGNTEXT"]:
        node_str += f'    <tag k="note" v="{props["SIGNTEXT"]}" />\n'
        
    # Extra OSM-helpful metadata indicating support structure types
    if "SUPPORTTYPE" in props and props["SUPPORTTYPE"]:
        node_str += f'    <tag k="support" v="{props["SUPPORTTYPE"].lower()}" />\n'
        
    node_str += f'    <tag k="source" v="City of Plano GIS" />\n'
    node_str += '  </node>\n'
    
    osm_xml += node_str

osm_xml += "\n</osm>\n"

with open("plano_stop_signs.osm", "w") as f:
    f.write(osm_xml)

print("Successfully wrote plano_stop_signs.osm")
