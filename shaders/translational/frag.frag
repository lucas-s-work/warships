#version 410
uniform sampler2D tex;
in vec2 fragtexcoord;

out vec4 frag_colour;

void main() {
    frag_colour = texture(tex, fragtexcoord);

    if(frag_colour.xyz == vec3(0,0,0)) {
        discard;
    }
}
