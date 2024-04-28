# Open311 Message Examples

**QUICK LINKS:** [Service List](###-Service-List)

## Example Responses

### Service List

**XML -EXAMPLE (e.g. https://api.city.gov/dev/v2/services.xml?jurisdiction_id=city.gov)**

```
<?xml version="1.0" encoding="utf-8"?>
<services>
    <service>
        <service_code>001</service_code>
        <service_name>Cans left out 24x7</service_name>
        <description>Garbage or recycling cans that have been left out for more than 24 hours after collection. Violators will be cited.</description>
        <metadata>true</metadata>
        <type>realtime</type>
        <keywords>lorem, ipsum, dolor</keywords>
        <group>sanitation</group>
    </service>
    <service>
        <service_code>002</service_code>
        <metadata>true</metadata>
        <type>realtime</type>
        <keywords>lorem, ipsum, dolor</keywords>
        <group>street</group>
        <service_name>Construction plate shifted</service_name>
        <description>Metal construction plate covering the street or sidewalk has been moved.</description>
    </service>
    <service>
        <service_code>003</service_code>
        <metadata>true</metadata>
        <type>realtime</type>
        <keywords>lorem, ipsum, dolor</keywords>
        <group>street</group>
        <service_name>Curb or curb ramp defect</service_name>
        <description>Sidewalk curb or ramp has problems such as cracking, missing pieces, holes, and/or chipped curb.</description>
    </service>
</services>
```

**JSON -EXAMPLE (e.g. https://api.city.gov/dev/v2/services.json?jurisdiction_id=city.gov)**

```
[
  {
    "service_code":001,
    "service_name":"Cans left out 24x7",
    "description":"Garbage or recycling cans that have been left out for more than 24 hours after collection. Violators will be cited.",
    "metadata":true,
    "type":"realtime",
    "keywords":"lorem, ipsum, dolor",
    "group":"sanitation"
  },
  {
    "service_code":002,
    "metadata":true,
    "type":"realtime",
    "keywords":"lorem, ipsum, dolor",
    "group":"street",
    "service_name":"Construction plate shifted",
    "description":"Metal construction plate covering the street or sidewalk has been moved."
  },
  {
    "service_code":003,
    "metadata":true,
    "type":"realtime",
    "keywords":"lorem, ipsum, dolor",
    "group":"street",
    "service_name":"Curb or curb ramp defect",
    "description":"Sidewalk curb or ramp has problems such as cracking, missing pieces, holes, and/or chipped curb."
  }
]
```

Source: https://wiki.open311.org/GeoReport_v2/ 