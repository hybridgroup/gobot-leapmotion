# Events

## hand

Gets triggered when doing a hand motion, makes hand input available.

Leap Hand contains:

- **position** (palm x-y-z coords)
- **wrist rotation** (axis,angle,matrix)

##### Example JSON

    

	  {
	      "direction": [
	        0.772435,
	        0.520335,
	        -0.364136
	      ],
	      "id": 57,
	      "palmNormal": [
	        -0.0100593,
	        -0.563263,
	        -0.826217
	      ],
	      "palmPosition": [
	        117.546,
	        236.007,
	        76.3394
	      ],
	      "palmVelocity": [
	        -866.196,
	        -100.749,
	        275.692
	      ],
	      "r": [
	        [
	          0.999844,
	          0.0142022,
	          0.0105289
	        ],
	        [
	          -0.0141201,
	          0.99987,
	          -0.00783186
	        ],
	        [
	          -0.0106388,
	          0.00768197,
	          0.999914
	        ]
	      ],
	      "s": 0.992511,
	      "sphereCenter": [
	        156.775,
	        227.378,
	        48.3453
	      ],
	      "sphereRadius": 75.3216,
	      "stabilizedPalmPosition": [
	        119.009,
	        236.071,
	        75.951
	      ],
	      "t": [
	        -38.0468,
	        28.2341,
	        -21.3291
	      ],
	      "timeVisible": 0.051952
	    }
  

## gesture

Gets triggered when doing a gesture motion, makes hand input available.

There is 4 Gesture types :

- **Circle** - A circular movement by a finger.
- **Swipe** - A straight line movement by the hand with fingers extended.
- **KeyTap** -  A downward tapping movement by a finger.
- **ScreenTap** - A forward tapping movement by a finger.

##### Example JSON

    
	  {
	      "direction": [
	        -0.647384,
	        0.750476,
	        -0.132964
	      ],
	      "duration": 0,
	      "handIds": [
	        57
	      ],
	      "id": 72,
	      "pointableIds": [
	        14
	      ],
	      "position": [
	        117.665,
	        313.471,
	        27.2095
	      ],
	      "speed": 1050.66,
	      "startPosition": [
	        195.438,
	        223.313,
	        43.183
	      ],
	      "state": "start",
	      "type": "swipe"
	    }


## pointable

Gets triggered when doing a pointable motion, makes hand input available.

Leap Pointable contains:

- **length** 
- **touch zone** 

##### Example JSON

    
	  {
	      "direction": [
	        0.54044,
	        0.174084,
	        -0.823176
	      ],
	      "handId": 57,
	      "id": 1,
	      "length": 48.393,
	      "stabilizedTipPosition": [
	        194.714,
	        291.812,
	        20.6219
	      ],
	      "timeVisible": 0.13873,
	      "tipPosition": [
	        194.714,
	        291.812,
	        20.6219
	      ],
	      "tipVelocity": [
	        -716.414,
	        686.468,
	        -427.914
	      ],
	      "tool": false,
	      "touchDistance": 0.333333,
	      "touchZone": "hovering"
	    }

## frame

Gets triggered with every motion detected, makes frame input available.

## Message

Gets triggered when receiving a message.

More information abour the parser [https://github.com/hybridgroup/gobot-leapmotion/blob/master/parser.go](https://github.com/hybridgroup/gobot-leapmotion/blob/master/parser.go).