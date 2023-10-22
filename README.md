![Logo](http://svg.wiersma.co.za/glasslabs/module?title=CLOCK&tag=a%20simple%20clock%20module)

Clock is a simple clock module for [looking glass](http://github.com/glasslabs/looking-glass)

![Screenshot](.readme/screenshot.png)

## Usage

Clone the clock into a path under your modules path and add the module path
to under modules in your configuration.

```yaml
modules:
 - name: simple-clock
    url:  https://github.com/glasslabs/clock/releases/download/v1.0.0/clock.wasm
    position: top:right
    config:
      timeFormat: 15:04
```

## Configuration

### Time Format (timeFormat)

*Default: 15:04*

Formats the time display of the clock using Go's [time formatting syntax](https://golang.org/pkg/time/#Time.Format).

### Date Format (dateFormat)

*Default: Monday, January 2*

Formats the date display of the clock using Go's [time formatting syntax](https://golang.org/pkg/time/#Time.Format).

### Time Zone (timezone)

*Default: Local*

The timezone name according to [IANA Time Zone databse](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)