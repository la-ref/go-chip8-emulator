#include <stdint.h>
#include <stdlib.h>

typedef struct {
    uint8_t volume;
    uint32_t squareWaveFreq;
    uint32_t sampleRate;
} AudioConfig;

void AudioCallback(void *userdata, uint8_t *stream, int len);