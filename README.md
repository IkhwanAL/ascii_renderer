# What is this Project

A Converter From Image to Ascii Image

## What is The Motivation of This Project

The plan is to learn Data Structure & Algorithm Based On Project. And how scaling image work. And doing something fun to be able to see image in terminal instead pixel this is character (The Wonder Of Math)

few image manipulation that i want learn that is
- Nearest Neighbor Interpolation [Done]
- Bilinear Interpolation [Done]
- Bicubic Interpolation
- Edge Detection [Done]
- Gaussian blur
- Dithering [Go already provide the algo]

After Few Weeks Of Building This Little Side/Fun Project There are few thing i learn
the word above like nearest neighbor or other it sound scary but with little of time it's not that difficult to implement (maybe part of it golang already provide the std package like image, color, etc)

## My notes
My First Implementation are Convert An Image To GrayScale (Already exists in golang) and Use That GrayScale And DownScale (Nearest Neighbor) so it Appropriate for terminal and Render to ascii, there already pre-sorted character that we can use in internet using that. And get the index based on their brightness of their gray color. print it. 

After completed it i increase the difficulty by using Bilinear Interpolation for image scale which instead of using 2 node in liniear interpolation we increase to 4 node it mean 2 linear interpolation (future me probaly thinking add other interpolation called bicubic interpolation which using 16 node (4 linear interpolation)) after that quick test the result sucess. 

After that i increase the difficult by adding edge detection, to be honest in edge detection this is the part that i need a lot of learning like what is kernel convolution, what is filter, etc. 

Which i manage successfully implemented the edge detection, but there's the problem down scale image mean losing pixel information and i can't get the edge result if image are losing color precision. so i should get edge result before downscale img, which invite another problem how do i render these two different size of image where the grayscale is downscaled but the edge result image is not. 

So i have to downscale the edge result image, try using bilinear failed, look at google, ask an AI, and the result is using MaxPooling another downscale just for edge detection image. another logic to create but eventually success. Now im confusing how the program know to render from those two image, after few hours, i should use threshold and mean to determine that if the edge is to thin render normally if not render with list of edge character. 

It Still Many Possiblity To Improve this Project which i don't know what algorirthm that i should use. And this project are lack of test(me just lazy create test that's all)

### Rant
for the Maxpooling part, that part ugh... its painful. soo many edge case to the point that i don't care anymore about the edge case, if the result are good, than it's good

Here some tips for you: to start a side project start from small and improve it little by little without realizing you already create some many thing

What do i learn from this all fun but painful:

Image Processing it's Hard especially debugging this. And you know the boundary, making sure they using a correct pixel if not BOOM... failed. especially the image 1000 x 1000 ohh boy good luck finding incorrect pixel in 1.000.000 pixel. And sometimes you don't know if the result are correct because it's produce the result but your is math is ways off. i have to use AI to make sure the math is correct because seeing soo many "+-/*" 

Well from all this rant i did quite enjoy this project, learn soo many thing, even forget what time is it. Usually when i code i just code for 1 - 2 hour, but since the project i build is intresting and fun without realizing it already 5+ hours.

### Algo
Step By Step Algo:

1. Get Image
2. Convert To GrayScale
3. Get Edge Detection image
4. Downscale The GrayScale Image
5. Downscale The Edge Detection Image
6. Render Image To Ascii

### Terms
New Terms for all search that i through:
- kernel convolution
- filter
- MaxPooling
- Nearest Neighbor
- Bilinear Interpolation
- Linear Interpolation
- Bicubic Interpolation
- Dithering
- Gaussian blur
- Edge Detection
- CNN (Convolution Neural Network)
