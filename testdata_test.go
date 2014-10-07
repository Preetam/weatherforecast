package weatherforecast

var testApiResponse = []byte(`
{
  "city": {
    "coord": {
      "lat": 38.029308,
      "lon": -78.476677
    },
    "country": "US",
    "id": 4752031,
    "name": "Charlottesville",
    "population": 0,
    "sys": {
      "population": 0
    }
  },
  "cnt": 37,
  "cod": "200",
  "list": [
    {
      "clouds": {
        "all": 36
      },
      "dt": 1412650800,
      "dt_txt": "2014-10-07 03:00:00",
      "main": {
        "grnd_level": 981.58,
        "humidity": 61,
        "pressure": 981.58,
        "sea_level": 1027.24,
        "temp": 289.4,
        "temp_kf": -1.1,
        "temp_max": 290.504,
        "temp_min": 289.4
      },
      "sys": {
        "pod": "n"
      },
      "weather": [
        {
          "description": "scattered clouds",
          "icon": "03n",
          "id": 802,
          "main": "Clouds"
        }
      ],
      "wind": {
        "deg": 201.501,
        "speed": 4.36
      }
    },
    {
      "clouds": {
        "all": 76
      },
      "dt": 1412661600,
      "dt_txt": "2014-10-07 06:00:00",
      "main": {
        "grnd_level": 981.63,
        "humidity": 61,
        "pressure": 981.63,
        "sea_level": 1027.68,
        "temp": 289.44,
        "temp_kf": -1.05,
        "temp_max": 290.486,
        "temp_min": 289.44
      },
      "sys": {
        "pod": "n"
      },
      "weather": [
        {
          "description": "broken clouds",
          "icon": "04n",
          "id": 803,
          "main": "Clouds"
        }
      ],
      "wind": {
        "deg": 211.01,
        "speed": 3.82
      }
    },
    {
      "clouds": {
        "all": 92
      },
      "dt": 1412672400,
      "dt_txt": "2014-10-07 09:00:00",
      "main": {
        "grnd_level": 983.36,
        "humidity": 86,
        "pressure": 983.36,
        "sea_level": 1029.24,
        "temp": 286.75,
        "temp_kf": -0.99,
        "temp_max": 287.745,
        "temp_min": 286.75
      },
      "sys": {
        "pod": "n"
      },
      "weather": [
        {
          "description": "overcast clouds",
          "icon": "04n",
          "id": 804,
          "main": "Clouds"
        }
      ],
      "wind": {
        "deg": 248.501,
        "speed": 2.62
      }
    },
    {
      "clouds": {
        "all": 80
      },
      "dt": 1412683200,
      "dt_txt": "2014-10-07 12:00:00",
      "main": {
        "grnd_level": 984.39,
        "humidity": 78,
        "pressure": 984.39,
        "sea_level": 1030.35,
        "temp": 286.35,
        "temp_kf": -0.94,
        "temp_max": 287.287,
        "temp_min": 286.35
      },
      "sys": {
        "pod": "d"
      },
      "weather": [
        {
          "description": "broken clouds",
          "icon": "04d",
          "id": 803,
          "main": "Clouds"
        }
      ],
      "wind": {
        "deg": 249.502,
        "speed": 2.42
      }
    },
    {
      "clouds": {
        "all": 24
      },
      "dt": 1412694000,
      "dt_txt": "2014-10-07 15:00:00",
      "main": {
        "grnd_level": 984.43,
        "humidity": 63,
        "pressure": 984.43,
        "sea_level": 1030.04,
        "temp": 291.22,
        "temp_kf": -0.88,
        "temp_max": 292.099,
        "temp_min": 291.22
      },
      "sys": {
        "pod": "d"
      },
      "weather": [
        {
          "description": "few clouds",
          "icon": "02d",
          "id": 801,
          "main": "Clouds"
        }
      ],
      "wind": {
        "deg": 214.003,
        "speed": 2.02
      }
    },
    {
      "clouds": {
        "all": 36
      },
      "dt": 1412704800,
      "dt_txt": "2014-10-07 18:00:00",
      "main": {
        "grnd_level": 981.72,
        "humidity": 55,
        "pressure": 981.72,
        "sea_level": 1027.24,
        "temp": 293.63,
        "temp_kf": -0.83,
        "temp_max": 294.458,
        "temp_min": 293.63
      },
      "sys": {
        "pod": "d"
      },
      "weather": [
        {
          "description": "scattered clouds",
          "icon": "03d",
          "id": 802,
          "main": "Clouds"
        }
      ],
      "wind": {
        "deg": 193.007,
        "speed": 3.06
      }
    },
    {
      "clouds": {
        "all": 80
      },
      "dt": 1412715600,
      "dt_txt": "2014-10-07 21:00:00",
      "main": {
        "grnd_level": 978.86,
        "humidity": 56,
        "pressure": 978.86,
        "sea_level": 1024.06,
        "temp": 293.11,
        "temp_kf": -0.77,
        "temp_max": 293.883,
        "temp_min": 293.11
      },
      "sys": {
        "pod": "d"
      },
      "weather": [
        {
          "description": "broken clouds",
          "icon": "04d",
          "id": 803,
          "main": "Clouds"
        }
      ],
      "wind": {
        "deg": 192.504,
        "speed": 4.32
      }
    },
    {
      "clouds": {
        "all": 100
      },
      "dt": 1412726400,
      "dt_txt": "2014-10-08 00:00:00",
      "main": {
        "grnd_level": 977.42,
        "humidity": 89,
        "pressure": 977.42,
        "sea_level": 1022.52,
        "temp": 288.71,
        "temp_kf": -0.72,
        "temp_max": 289.424,
        "temp_min": 288.71
      },
      "rain": {
        "3h": 2
      },
      "sys": {
        "pod": "n"
      },
      "weather": [
        {
          "description": "light rain",
          "icon": "10n",
          "id": 500,
          "main": "Rain"
        }
      ],
      "wind": {
        "deg": 205.004,
        "speed": 5.77
      }
    },
    {
      "clouds": {
        "all": 92
      },
      "dt": 1412737200,
      "dt_txt": "2014-10-08 03:00:00",
      "main": {
        "grnd_level": 978.57,
        "humidity": 98,
        "pressure": 978.57,
        "sea_level": 1023.88,
        "temp": 287.07,
        "temp_kf": -0.66,
        "temp_max": 287.736,
        "temp_min": 287.07
      },
      "rain": {
        "3h": 5
      },
      "sys": {
        "pod": "n"
      },
      "weather": [
        {
          "description": "moderate rain",
          "icon": "10n",
          "id": 501,
          "main": "Rain"
        }
      ],
      "wind": {
        "deg": 232.502,
        "speed": 3.8
      }
    },
    {
      "clouds": {
        "all": 24
      },
      "dt": 1412748000,
      "dt_txt": "2014-10-08 06:00:00",
      "main": {
        "grnd_level": 979.3,
        "humidity": 86,
        "pressure": 979.3,
        "sea_level": 1025.08,
        "temp": 286.19,
        "temp_kf": -0.61,
        "temp_max": 286.798,
        "temp_min": 286.19
      },
      "rain": {
        "3h": 2
      },
      "sys": {
        "pod": "n"
      },
      "weather": [
        {
          "description": "light rain",
          "icon": "10n",
          "id": 500,
          "main": "Rain"
        }
      ],
      "wind": {
        "deg": 261.007,
        "speed": 3.3
      }
    },
    {
      "clouds": {
        "all": 32
      },
      "dt": 1412758800,
      "dt_txt": "2014-10-08 09:00:00",
      "main": {
        "grnd_level": 980.29,
        "humidity": 83,
        "pressure": 980.29,
        "sea_level": 1026.34,
        "temp": 285.08,
        "temp_kf": -0.55,
        "temp_max": 285.632,
        "temp_min": 285.08
      },
      "rain": {
        "3h": 0
      },
      "sys": {
        "pod": "n"
      },
      "weather": [
        {
          "description": "scattered clouds",
          "icon": "03n",
          "id": 802,
          "main": "Clouds"
        }
      ],
      "wind": {
        "deg": 246.001,
        "speed": 2.3
      }
    },
    {
      "clouds": {
        "all": 56
      },
      "dt": 1412769600,
      "dt_txt": "2014-10-08 12:00:00",
      "main": {
        "grnd_level": 982.58,
        "humidity": 85,
        "pressure": 982.58,
        "sea_level": 1028.83,
        "temp": 285.09,
        "temp_kf": -0.5,
        "temp_max": 285.582,
        "temp_min": 285.09
      },
      "rain": {
        "3h": 0
      },
      "sys": {
        "pod": "d"
      },
      "weather": [
        {
          "description": "broken clouds",
          "icon": "04d",
          "id": 803,
          "main": "Clouds"
        }
      ],
      "wind": {
        "deg": 282,
        "speed": 1.41
      }
    },
    {
      "clouds": {
        "all": 12
      },
      "dt": 1412780400,
      "dt_txt": "2014-10-08 15:00:00",
      "main": {
        "grnd_level": 983.5,
        "humidity": 71,
        "pressure": 983.5,
        "sea_level": 1029.41,
        "temp": 288.98,
        "temp_kf": -0.44,
        "temp_max": 289.423,
        "temp_min": 288.98
      },
      "rain": {
        "3h": 0
      },
      "sys": {
        "pod": "d"
      },
      "weather": [
        {
          "description": "few clouds",
          "icon": "02d",
          "id": 801,
          "main": "Clouds"
        }
      ],
      "wind": {
        "deg": 283.006,
        "speed": 1.91
      }
    },
    {
      "clouds": {
        "all": 0
      },
      "dt": 1412791200,
      "dt_txt": "2014-10-08 18:00:00",
      "main": {
        "grnd_level": 983.04,
        "humidity": 49,
        "pressure": 983.04,
        "sea_level": 1028.73,
        "temp": 291.7,
        "temp_kf": -0.39,
        "temp_max": 292.082,
        "temp_min": 291.7
      },
      "rain": {
        "3h": 0
      },
      "sys": {
        "pod": "d"
      },
      "weather": [
        {
          "description": "sky is clear",
          "icon": "01d",
          "id": 800,
          "main": "Clear"
        }
      ],
      "wind": {
        "deg": 284.502,
        "speed": 3.17
      }
    },
    {
      "clouds": {
        "all": 0
      },
      "dt": 1412802000,
      "dt_txt": "2014-10-08 21:00:00",
      "main": {
        "grnd_level": 983.28,
        "humidity": 45,
        "pressure": 983.28,
        "sea_level": 1028.77,
        "temp": 291.4,
        "temp_kf": -0.33,
        "temp_max": 291.735,
        "temp_min": 291.4
      },
      "rain": {
        "3h": 0
      },
      "sys": {
        "pod": "d"
      },
      "weather": [
        {
          "description": "sky is clear",
          "icon": "01d",
          "id": 800,
          "main": "Clear"
        }
      ],
      "wind": {
        "deg": 288.502,
        "speed": 3.11
      }
    },
    {
      "clouds": {
        "all": 0
      },
      "dt": 1412812800,
      "dt_txt": "2014-10-09 00:00:00",
      "main": {
        "grnd_level": 984.87,
        "humidity": 51,
        "pressure": 984.87,
        "sea_level": 1030.84,
        "temp": 286.36,
        "temp_kf": -0.28,
        "temp_max": 286.639,
        "temp_min": 286.36
      },
      "rain": {
        "3h": 0
      },
      "sys": {
        "pod": "n"
      },
      "weather": [
        {
          "description": "sky is clear",
          "icon": "01n",
          "id": 800,
          "main": "Clear"
        }
      ],
      "wind": {
        "deg": 293.501,
        "speed": 2.01
      }
    },
    {
      "clouds": {
        "all": 0
      },
      "dt": 1412823600,
      "dt_txt": "2014-10-09 03:00:00",
      "main": {
        "grnd_level": 986.45,
        "humidity": 65,
        "pressure": 986.45,
        "sea_level": 1032.89,
        "temp": 282.06,
        "temp_kf": -0.22,
        "temp_max": 282.284,
        "temp_min": 282.06
      },
      "rain": {
        "3h": 0
      },
      "sys": {
        "pod": "n"
      },
      "weather": [
        {
          "description": "sky is clear",
          "icon": "01n",
          "id": 800,
          "main": "Clear"
        }
      ],
      "wind": {
        "deg": 286.502,
        "speed": 1.61
      }
    },
    {
      "clouds": {
        "all": 0
      },
      "dt": 1412834400,
      "dt_txt": "2014-10-09 06:00:00",
      "main": {
        "grnd_level": 986.92,
        "humidity": 83,
        "pressure": 986.92,
        "sea_level": 1033.59,
        "temp": 279.12,
        "temp_kf": -0.17,
        "temp_max": 279.284,
        "temp_min": 279.12
      },
      "rain": {
        "3h": 0
      },
      "sys": {
        "pod": "n"
      },
      "weather": [
        {
          "description": "sky is clear",
          "icon": "01n",
          "id": 800,
          "main": "Clear"
        }
      ],
      "wind": {
        "deg": 271.012,
        "speed": 0.96
      }
    },
    {
      "clouds": {
        "all": 0
      },
      "dt": 1412845200,
      "dt_txt": "2014-10-09 09:00:00",
      "main": {
        "grnd_level": 986.94,
        "humidity": 83,
        "pressure": 986.94,
        "sea_level": 1033.73,
        "temp": 278.09,
        "temp_kf": -0.11,
        "temp_max": 278.204,
        "temp_min": 278.09
      },
      "rain": {
        "3h": 0
      },
      "sys": {
        "pod": "n"
      },
      "weather": [
        {
          "description": "sky is clear",
          "icon": "01n",
          "id": 800,
          "main": "Clear"
        }
      ],
      "wind": {
        "deg": 208.001,
        "speed": 0.97
      }
    },
    {
      "clouds": {
        "all": 20
      },
      "dt": 1412856000,
      "dt_txt": "2014-10-09 12:00:00",
      "main": {
        "grnd_level": 987.38,
        "humidity": 83,
        "pressure": 987.38,
        "sea_level": 1034.17,
        "temp": 278.77,
        "temp_kf": -0.06,
        "temp_max": 278.828,
        "temp_min": 278.77
      },
      "rain": {
        "3h": 0
      },
      "sys": {
        "pod": "d"
      },
      "weather": [
        {
          "description": "few clouds",
          "icon": "02d",
          "id": 801,
          "main": "Clouds"
        }
      ],
      "wind": {
        "deg": 170.502,
        "speed": 0.97
      }
    },
    {
      "clouds": {
        "all": 48
      },
      "dt": 1412866800,
      "dt_txt": "2014-10-09 15:00:00",
      "main": {
        "grnd_level": 987.4,
        "humidity": 54,
        "pressure": 987.4,
        "sea_level": 1033.43,
        "temp": 289.767,
        "temp_max": 289.767,
        "temp_min": 289.767
      },
      "rain": {
        "3h": 0
      },
      "sys": {
        "pod": "d"
      },
      "weather": [
        {
          "description": "scattered clouds",
          "icon": "03d",
          "id": 802,
          "main": "Clouds"
        }
      ],
      "wind": {
        "deg": 165.003,
        "speed": 1.36
      }
    },
    {
      "clouds": {
        "all": 76
      },
      "dt": 1412877600,
      "dt_txt": "2014-10-09 18:00:00",
      "main": {
        "grnd_level": 985.05,
        "humidity": 53,
        "pressure": 985.05,
        "sea_level": 1030.63,
        "temp": 291.286,
        "temp_max": 291.286,
        "temp_min": 291.286
      },
      "rain": {
        "3h": 0
      },
      "sys": {
        "pod": "d"
      },
      "weather": [
        {
          "description": "broken clouds",
          "icon": "04d",
          "id": 803,
          "main": "Clouds"
        }
      ],
      "wind": {
        "deg": 195.003,
        "speed": 2.26
      }
    },
    {
      "clouds": {
        "all": 76
      },
      "dt": 1412888400,
      "dt_txt": "2014-10-09 21:00:00",
      "main": {
        "grnd_level": 982.12,
        "humidity": 56,
        "pressure": 982.12,
        "sea_level": 1027.55,
        "temp": 292.663,
        "temp_max": 292.663,
        "temp_min": 292.663
      },
      "rain": {
        "3h": 0
      },
      "sys": {
        "pod": "d"
      },
      "weather": [
        {
          "description": "broken clouds",
          "icon": "04d",
          "id": 803,
          "main": "Clouds"
        }
      ],
      "wind": {
        "deg": 201,
        "speed": 3.02
      }
    },
    {
      "clouds": {
        "all": 92
      },
      "dt": 1412899200,
      "dt_txt": "2014-10-10 00:00:00",
      "main": {
        "grnd_level": 980.87,
        "humidity": 60,
        "pressure": 980.87,
        "sea_level": 1026.48,
        "temp": 291.932,
        "temp_max": 291.932,
        "temp_min": 291.932
      },
      "rain": {
        "3h": 0
      },
      "sys": {
        "pod": "n"
      },
      "weather": [
        {
          "description": "overcast clouds",
          "icon": "04n",
          "id": 804,
          "main": "Clouds"
        }
      ],
      "wind": {
        "deg": 205.5,
        "speed": 3.66
      }
    },
    {
      "clouds": {
        "all": 80
      },
      "dt": 1412910000,
      "dt_txt": "2014-10-10 03:00:00",
      "main": {
        "grnd_level": 980.57,
        "humidity": 65,
        "pressure": 980.57,
        "sea_level": 1026.03,
        "temp": 292.864,
        "temp_max": 292.864,
        "temp_min": 292.864
      },
      "rain": {
        "3h": 0
      },
      "sys": {
        "pod": "n"
      },
      "weather": [
        {
          "description": "broken clouds",
          "icon": "04n",
          "id": 803,
          "main": "Clouds"
        }
      ],
      "wind": {
        "deg": 224.5,
        "speed": 3.46
      }
    },
    {
      "clouds": {
        "all": 56
      },
      "dt": 1412920800,
      "dt_txt": "2014-10-10 06:00:00",
      "main": {
        "grnd_level": 979.09,
        "humidity": 72,
        "pressure": 979.09,
        "sea_level": 1024.6,
        "temp": 293.249,
        "temp_max": 293.249,
        "temp_min": 293.249
      },
      "rain": {
        "3h": 0
      },
      "sys": {
        "pod": "n"
      },
      "weather": [
        {
          "description": "broken clouds",
          "icon": "04n",
          "id": 803,
          "main": "Clouds"
        }
      ],
      "wind": {
        "deg": 240.002,
        "speed": 2.92
      }
    },
    {
      "clouds": {
        "all": 32
      },
      "dt": 1412931600,
      "dt_txt": "2014-10-10 09:00:00",
      "main": {
        "grnd_level": 978.55,
        "humidity": 79,
        "pressure": 978.55,
        "sea_level": 1023.85,
        "temp": 292.751,
        "temp_max": 292.751,
        "temp_min": 292.751
      },
      "rain": {
        "3h": 0
      },
      "sys": {
        "pod": "n"
      },
      "weather": [
        {
          "description": "scattered clouds",
          "icon": "03n",
          "id": 802,
          "main": "Clouds"
        }
      ],
      "wind": {
        "deg": 255.002,
        "speed": 2.31
      }
    },
    {
      "clouds": {
        "all": 48
      },
      "dt": 1412942400,
      "dt_txt": "2014-10-10 12:00:00",
      "main": {
        "grnd_level": 979.35,
        "humidity": 86,
        "pressure": 979.35,
        "sea_level": 1024.86,
        "temp": 291.104,
        "temp_max": 291.104,
        "temp_min": 291.104
      },
      "rain": {
        "3h": 0
      },
      "sys": {
        "pod": "d"
      },
      "weather": [
        {
          "description": "scattered clouds",
          "icon": "03d",
          "id": 802,
          "main": "Clouds"
        }
      ],
      "wind": {
        "deg": 263.501,
        "speed": 1.36
      }
    },
    {
      "clouds": {
        "all": 48
      },
      "dt": 1412953200,
      "dt_txt": "2014-10-10 15:00:00",
      "main": {
        "grnd_level": 979.86,
        "humidity": 64,
        "pressure": 979.86,
        "sea_level": 1025.13,
        "temp": 296.269,
        "temp_max": 296.269,
        "temp_min": 296.269
      },
      "rain": {
        "3h": 0
      },
      "sys": {
        "pod": "d"
      },
      "weather": [
        {
          "description": "scattered clouds",
          "icon": "03d",
          "id": 802,
          "main": "Clouds"
        }
      ],
      "wind": {
        "deg": 261,
        "speed": 1.42
      }
    },
    {
      "clouds": {
        "all": 92
      },
      "dt": 1412964000,
      "dt_txt": "2014-10-10 18:00:00",
      "main": {
        "grnd_level": 979.29,
        "humidity": 77,
        "pressure": 979.29,
        "sea_level": 1024.45,
        "temp": 294.241,
        "temp_max": 294.241,
        "temp_min": 294.241
      },
      "rain": {
        "3h": 0
      },
      "sys": {
        "pod": "d"
      },
      "weather": [
        {
          "description": "overcast clouds",
          "icon": "04d",
          "id": 804,
          "main": "Clouds"
        }
      ],
      "wind": {
        "deg": 239.005,
        "speed": 1.41
      }
    },
    {
      "clouds": {
        "all": 48
      },
      "dt": 1412974800,
      "dt_txt": "2014-10-10 21:00:00",
      "main": {
        "grnd_level": 978.97,
        "humidity": 73,
        "pressure": 978.97,
        "sea_level": 1024.49,
        "temp": 294.695,
        "temp_max": 294.695,
        "temp_min": 294.695
      },
      "rain": {
        "3h": 0
      },
      "sys": {
        "pod": "d"
      },
      "weather": [
        {
          "description": "scattered clouds",
          "icon": "03d",
          "id": 802,
          "main": "Clouds"
        }
      ],
      "wind": {
        "deg": 276.002,
        "speed": 2.1
      }
    },
    {
      "clouds": {
        "all": 0
      },
      "dt": 1412996400,
      "dt_txt": "2014-10-11 03:00:00",
      "main": {
        "grnd_level": 984.22,
        "humidity": 68,
        "pressure": 984.22,
        "sea_level": 1030.37,
        "temp": 286.436,
        "temp_max": 286.436,
        "temp_min": 286.436
      },
      "rain": {
        "3h": 0
      },
      "sys": {
        "pod": "n"
      },
      "weather": [
        {
          "description": "sky is clear",
          "icon": "01n",
          "id": 800,
          "main": "Clear"
        }
      ],
      "wind": {
        "deg": 333,
        "speed": 2.57
      }
    },
    {
      "clouds": {
        "all": 8
      },
      "dt": 1413007200,
      "dt_txt": "2014-10-11 06:00:00",
      "main": {
        "grnd_level": 985.63,
        "humidity": 65,
        "pressure": 985.63,
        "sea_level": 1032.41,
        "temp": 283.103,
        "temp_max": 283.103,
        "temp_min": 283.103
      },
      "rain": {
        "3h": 0
      },
      "sys": {
        "pod": "n"
      },
      "weather": [
        {
          "description": "sky is clear",
          "icon": "02n",
          "id": 800,
          "main": "Clear"
        }
      ],
      "wind": {
        "deg": 339.504,
        "speed": 2.01
      }
    },
    {
      "clouds": {
        "all": 12
      },
      "dt": 1413028800,
      "dt_txt": "2014-10-11 12:00:00",
      "main": {
        "grnd_level": 989.01,
        "humidity": 79,
        "pressure": 989.01,
        "sea_level": 1036.49,
        "temp": 277.97,
        "temp_max": 277.97,
        "temp_min": 277.97
      },
      "rain": {
        "3h": 0
      },
      "sys": {
        "pod": "d"
      },
      "weather": [
        {
          "description": "few clouds",
          "icon": "02d",
          "id": 801,
          "main": "Clouds"
        }
      ],
      "wind": {
        "deg": 344.507,
        "speed": 1.2
      }
    },
    {
      "clouds": {
        "all": 0
      },
      "dt": 1413039600,
      "dt_txt": "2014-10-11 15:00:00",
      "main": {
        "grnd_level": 990.31,
        "humidity": 52,
        "pressure": 990.31,
        "sea_level": 1037.18,
        "temp": 286.539,
        "temp_max": 286.539,
        "temp_min": 286.539
      },
      "rain": {
        "3h": 0
      },
      "sys": {
        "pod": "d"
      },
      "weather": [
        {
          "description": "sky is clear",
          "icon": "01d",
          "id": 800,
          "main": "Clear"
        }
      ],
      "wind": {
        "deg": 13.5013,
        "speed": 1.42
      }
    },
    {
      "clouds": {
        "all": 0
      },
      "dt": 1413050400,
      "dt_txt": "2014-10-11 18:00:00",
      "main": {
        "grnd_level": 989.08,
        "humidity": 46,
        "pressure": 989.08,
        "sea_level": 1035.51,
        "temp": 289.6,
        "temp_max": 289.6,
        "temp_min": 289.6
      },
      "rain": {
        "3h": 0
      },
      "sys": {
        "pod": "d"
      },
      "weather": [
        {
          "description": "sky is clear",
          "icon": "01d",
          "id": 800,
          "main": "Clear"
        }
      ],
      "wind": {
        "deg": 76.5028,
        "speed": 1.42
      }
    },
    {
      "clouds": {
        "all": 0
      },
      "dt": 1413061200,
      "dt_txt": "2014-10-11 21:00:00",
      "main": {
        "grnd_level": 988.43,
        "humidity": 46,
        "pressure": 988.43,
        "sea_level": 1034.71,
        "temp": 289.823,
        "temp_max": 289.823,
        "temp_min": 289.823
      },
      "rain": {
        "3h": 0
      },
      "sys": {
        "pod": "d"
      },
      "weather": [
        {
          "description": "sky is clear",
          "icon": "01d",
          "id": 800,
          "main": "Clear"
        }
      ],
      "wind": {
        "deg": 152.001,
        "speed": 1.36
      }
    }
  ],
  "message": 0.099
}
`)
