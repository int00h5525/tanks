from django.db import models
from django.core.validators import MinValueValidator, \
        MaxValueValidator, \
        RegexValidator

class Tank(models.Model):
    name = models.CharField(max_length=64)
    author = models.CharField(max_length=64)
    color = models.CharField(max_length=7, validators=[RegexValidator("#[0-9a-fA-F]{6}")])
    program = models.CharField(max_length=16384)

    radius_validators = [MinValueValidator(0), MaxValueValidator(100)]
    angle_validators = [MinValueValidator(-359), MaxValueValidator(359)]
    width_validators = [MinValueValidator(0), MaxValueValidator(100)]

    s0r = models.IntegerField(validators=radius_validators)
    s0a = models.IntegerField(validators=angle_validators)
    s0w = models.IntegerField(validators=width_validators)

    s1r = models.IntegerField(validators=radius_validators)
    s1a = models.IntegerField(validators=angle_validators)
    s1w = models.IntegerField(validators=width_validators)

    s2r = models.IntegerField(validators=radius_validators)
    s2a = models.IntegerField(validators=angle_validators)
    s2w = models.IntegerField(validators=width_validators)

    s3r = models.IntegerField(validators=radius_validators)
    s3a = models.IntegerField(validators=angle_validators)
    s3w = models.IntegerField(validators=width_validators)

    s4r = models.IntegerField(validators=radius_validators)
    s4a = models.IntegerField(validators=angle_validators)
    s4w = models.IntegerField(validators=width_validators)

    s5r = models.IntegerField(validators=radius_validators)
    s5a = models.IntegerField(validators=angle_validators)
    s5w = models.IntegerField(validators=width_validators)

    s6r = models.IntegerField(validators=radius_validators)
    s6a = models.IntegerField(validators=angle_validators)
    s6w = models.IntegerField(validators=width_validators)

    s7r = models.IntegerField(validators=radius_validators)
    s7a = models.IntegerField(validators=angle_validators)
    s7w = models.IntegerField(validators=width_validators)

    s8r = models.IntegerField(validators=radius_validators)
    s8a = models.IntegerField(validators=angle_validators)
    s8w = models.IntegerField(validators=width_validators)

    s9r = models.IntegerField(validators=radius_validators)
    s9a = models.IntegerField(validators=angle_validators)
    s9w = models.IntegerField(validators=width_validators)
