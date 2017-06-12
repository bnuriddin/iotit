package repo

const (
	file = "mapping.json"

	// example is used to create mapper.json file or used in case of absence of the file
	example = `{
  "Devices": [
    {
      "Name": "raspberry-pi",
      "Sub": [
        {
          "Name": "Raspberry Pi Model A, A+, B, B+, Zero, Zero W"
        },
        {
          "Name": "Raspberry Pi 2 (based on Model B+)"
        },
        {
          "Name": "Raspberry Pi 3"
        }
      ],
      "Images": [
        {
          "URL": "http://director.downloads.raspberrypi.org/raspbian_lite/images/raspbian_lite-2017-03-03/2017-03-02-raspbian-jessie-lite.zip",
          "Title": "Raspbian jessie light",
          "User": "pi",
          "Pass": "raspberry"
        },
        {
          "URL": "http://director.downloads.raspberrypi.org/raspbian/images/raspbian-2017-03-03/2017-03-02-raspbian-jessie.zip",
          "Title": "Raspbian jessie w Pixel",
          "User": "pi",
          "Pass": "raspberry"
        }
      ]
    },
    {
      "Name": "edison",
      "Images": [
        {
          "URL": "http://iotdk.intel.com/images/3.5/edison/iot-devkit-prof-dev-image-edison-20160606.zip",
          "User": "root"
        }
      ]
    },
    {
      "Name": "nano-pi",
      "Sub": [
        {
          "Name": "NanoPi 2"
        },
        {
          "Name": "NanoPi 2 Fire"
        },
        {
          "Name": "NanoPi M1 Plus",
          "Images": [
            {
              "URL": "http://www.mediafire.com/file/cn292j2nxpz6j69/nanopi-m1-plus_ubuntu-core-xenial_4.11.2_20170531.img.zip",
              "Title": "Light Ubuntu-core",
              "User": "root",
              "Pass": "fa"
            },
            {
              "URL": "http://www.mediafire.com/file/ddquv7vfrkucb27/nanopi-m1-plus_debian-jessie_4.11.2_20170531.img.zip",
              "Title": "Debian Jessie",
              "User": "root",
              "Pass": "fa"
            }
          ]
        },
        {
          "Name": "NanoPi M1",
          "Images": [
            {
              "URL": "http://www.mediafire.com/file/kfihz987rf6s87d/nanopi-m1-debian-sd4g-20170204.img.zip",
              "Title": "Light Ubuntu-core",
              "User": "root",
              "Pass": "fa"
            },
            {
              "URL": "http://www.mediafire.com/file/kfihz987rf6s87d/nanopi-m1-debian-sd4g-20170204.img.zip",
              "Title": "Debian",
              "User": "root",
              "Pass": "fa"
            }
          ]
        },
        {
          "Name": "NanoPi M2"
        },
        {
          "Name": "NanoPi M3"
        },
        {
          "Name": "NanoPi NEO",
          "Images": [
            {
              "URL": "http://www.mediafire.com/file/hysi686cq89d9ll/nanopi-neo_ubuntu-core-xenial_4.11.2_20170605.img.zip",
              "User": "root",
              "Pass": "fa"
            }
          ]
        },
        {
          "Name": "NanoPi NEO 2",
          "Images": [
            {
              "URL": "http://www.mediafire.com/file/w2sp6pdwlj5l66j/nanopi-neo2_ubuntu-core-xenial_4.11.2_20170605.img.zip",
              "User": "root",
              "Pass": "fa"
            }
          ]
        },
        {
          "Name": "NanoPi NEO Air",
          "Images": [
            {
              "URL": "http://www.mediafire.com/file/i4jq7fqfp55bqhh/nanopi-neo-air_ubuntu-core-xenial_4.11.2_20170605.img.zip",
              "User": "root",
              "Pass": "fa"
            }
          ]
        },
        {
          "Name": "NanoPi S2"
        },
        {
          "Name": "NanoPi a64",
          "Images": [
            {
              "URL": "http://www.mediafire.com/file/hzfhcu0r1bb9ogt/nanopi-a64-core-qte-sd4g-20161129.img.zip",
              "Title": "Light Ubuntu-core",
              "User": "root",
              "Pass": "fa"
            },
            {
              "URL": "http://www.mediafire.com/file/cbh9mkb70p18m12/nanopi-a64-ubuntu-mate-sd4g-20161129.img.zip",
              "Title": "Ubuntu with a MATE-desktop",
              "User": "root",
              "Pass": "fa"
            }
          ]
        }
      ],
      "Images": [
        {
          "URL": "http://www.mediafire.com/file/q4cpxn4dmoo63r4/s5p4418-ubuntu-core-qte-sd4g-20170414.img.zip",
          "Title": "Light Ubuntu-core",
          "User": "root",
          "Pass": "fa"
        },
        {
          "URL": "http://www.mediafire.com/file/k1ghutjrecepqiv/s5p4418-debian-sd4g-20170414.img.zip",
          "Title": "Debian Jessie",
          "User": "root",
          "Pass": "fa"
        },
        {
          "URL": "http://www.mediafire.com/file/4e2es8j08vc8zp2/s5p4418-android-sd4g-20170414.img.zip",
          "Title": "Android"
        },
        {
          "URL": "http://www.mediafire.com/file/4293ylyitgmo4l8/s5p4418-debian-wifiap-sd4g-20170414.img.zip",
          "Title": "Debian WiFi AP",
          "User": "root",
          "Pass": "fa"
        }
      ]
    },
    {
      "Name": "beaglebone",
      "Sub": [
        {
          "Name": "BeagleBone (*,*Black,*Blue,*Green,*Wireless)"
        },
        {
          "Name": "BeagleBoard-X15",
          "Images": [
            {
              "URL": "https://debian.beagleboard.org/images/bbx15-debian-8.6-lxqt-4gb-armhf-2016-11-06-4gb.img.xz",
              "User": "ubuntu",
              "Pass": "temppwd"
            }
          ]
        },
        {
          "Name": "BeagleBoard-xM",
          "Images": [
            {
              "URL": "https://debian.beagleboard.org/images/bbxm-debian-8.6-lxqt-xm-4gb-armhf-2016-11-06-4gb.img.xz",
              "User": "ubuntu",
              "Pass": "temppwd"
            }
          ]
        }
      ],
      "Images": [
        {
          "URL": "https://cdn.isaax.io/isaax-distro/boards/beaglebone/ubuntu-16.04.1/bone-ubuntu-16.04.1-console-armhf-2016-10-06-2gb.zip",
          "Title": "Debian light isaax-distro",
          "User": "ubuntu",
          "Pass": "temppwd"
        },
        {
          "URL": "https://debian.beagleboard.org/images/bone-debian-8.7-lxqt-4gb-armhf-2017-03-19-4gb.img.xz",
          "Title": "Debian with desktop",
          "User": "ubuntu",
          "Pass": "temppwd"
        }
      ]
    }
  ]
}
`
)
