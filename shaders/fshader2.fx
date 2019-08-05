#version 330 core
out vec4 FragColor;

uniform float greenVal;

void main()
{
	FragColor = vec4(0.0, greenVal, 0.0, 0.0);
}
